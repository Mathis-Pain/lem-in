let turns = [];
let currentTurnIndex = 0;
let animationInterval = null;
let isPaused = false;
let speed = 1;
let rooms = {};
let paths = [];
let antsPositions = {};
let totalAnts = 0;
let finishedAnts = 0;
let finishedAntsSet = new Set();
let antPreviousPositions = {}; // Track previous positions for path following

window.onload = async function () {
  console.log("🔵 window.onload - Démarrage du chargement");
  try {
    const roomsResponse = await fetch("/api/rooms");
    const roomsData = await roomsResponse.json();
    console.log("✅ Rooms chargées:", Object.keys(roomsData).length, "salles");

    const movesResponse = await fetch("/api/moves");
    const movesDataResponse = await movesResponse.json();
    console.log("✅ Moves reçues:", movesDataResponse);

    const linksResponse = await fetch("/api/links");
    const linksData = await linksResponse.json();
    console.log("✅ Links chargés:", linksData.length, "liens");

    if (movesDataResponse.moves) {
      turns = movesDataResponse.moves
        .split("\n")
        .filter((line) => line.trim());
      
      console.log("✅ Turns parsés:", turns.length, "tours");
      console.log("Premier tour:", turns[0]);

      if (roomsData && Object.keys(roomsData).length > 0) {
        loadRoomsWithCoordinates(roomsData, linksData);
      }

      createGraph();
      document.querySelector(".loading").style.display = "none";
      document.getElementById("totalTurns").textContent = turns.length;
      console.log("✅ Initialisation terminée");
    } else {
      console.error("❌ Aucun mouvement reçu!");
    }
  } catch (error) {
    console.error("❌ Erreur de chargement:", error);
    document.querySelector(".loading").textContent =
      "Erreur: " + error.message;
  }
};

document.getElementById("speedRange").addEventListener("input", (e) => {
  speed = parseFloat(e.target.value);
  document.getElementById("speedValue").textContent =
    speed.toFixed(1) + "x";
  if (animationInterval) {
    clearInterval(animationInterval);
    animationInterval = setInterval(nextTurn, 1000 / speed);
  }
});

function loadRoomsWithCoordinates(roomsData, linksData) {
  rooms = {};
  paths = [];
  antsPositions = {};
  totalAnts = 0;
  finishedAnts = 0;

  let minX = Infinity,
    maxX = -Infinity;
  let minY = Infinity,
    maxY = -Infinity;

  Object.values(roomsData).forEach((room) => {
    if (room.x < minX) minX = room.x;
    if (room.x > maxX) maxX = room.x;
    if (room.y < minY) minY = room.y;
    if (room.y > maxY) maxY = room.y;
  });

  const svg = document.getElementById("graphSvg");
  const width = svg.clientWidth;
  const height = svg.clientHeight;
  const padding = 80;

  Object.entries(roomsData).forEach(([name, room]) => {
    const normalizedX =
      padding +
      ((room.x - minX) / (maxX - minX || 1)) * (width - 2 * padding);
    const normalizedY =
      padding +
      ((room.y - minY) / (maxY - minY || 1)) * (height - 2 * padding);

    rooms[name] = {
      x: normalizedX,
      y: normalizedY,
      name: name,
      type: room.type,
    };
  });

  linksData.forEach((link) => {
    if (rooms[link.from] && rooms[link.to]) {
      paths.push({ from: link.from, to: link.to });
    }
  });

  turns.forEach((turn) => {
    const moves = turn.split(" ");
    moves.forEach((move) => {
      const [ant, room] = move.split("-");
      const antNum = parseInt(ant.substring(1));
      if (antNum > totalAnts) totalAnts = antNum;
    });
  });

  document.getElementById("totalTurns").textContent = turns.length;
}

function createGraph() {
  const svg = document.getElementById("graphSvg");
  svg.innerHTML = "";

  paths.forEach((path) => {
    const from = rooms[path.from];
    const to = rooms[path.to];
    const line = document.createElementNS(
      "http://www.w3.org/2000/svg",
      "line"
    );
    line.setAttribute("x1", from.x);
    line.setAttribute("y1", from.y);
    line.setAttribute("x2", to.x);
    line.setAttribute("y2", to.y);
    line.setAttribute("class", "path");
    svg.appendChild(line);
  });

  Object.values(rooms).forEach((room) => {
    const g = document.createElementNS("http://www.w3.org/2000/svg", "g");
    g.setAttribute("class", `room ${room.type}`);

    const circle = document.createElementNS(
      "http://www.w3.org/2000/svg",
      "circle"
    );
    circle.setAttribute("cx", room.x);
    circle.setAttribute("cy", room.y);
    circle.setAttribute("r", 25);
    g.appendChild(circle);

    if (room.type === "start" || room.type === "end") {
      const holeInner = document.createElementNS(
        "http://www.w3.org/2000/svg",
        "circle"
      );
      holeInner.setAttribute("cx", room.x);
      holeInner.setAttribute("cy", room.y);
      holeInner.setAttribute("r", 15);
      holeInner.setAttribute("class", "hole-inner");
      g.appendChild(holeInner);
    }

    const text = document.createElementNS(
      "http://www.w3.org/2000/svg",
      "text"
    );
    text.setAttribute("x", room.x);
    text.setAttribute("y", room.y + 45);
    text.setAttribute("text-anchor", "middle");
    text.textContent = room.name;
    g.appendChild(text);

    svg.appendChild(g);
  });
}

function startAnimation() {
  console.log("🟢 startAnimation appelée");
  console.log("Nombre de tours disponibles:", turns.length);
  
  currentTurnIndex = 0;
  finishedAnts = 0;
  finishedAntsSet.clear();
  antPreviousPositions = {};
  isPaused = false;
  document.getElementById("currentTurn").textContent = "0";
  document.getElementById("activeAnts").textContent = "0";
  document.getElementById("finishedAnts").textContent = "0";

  const svg = document.getElementById("graphSvg");
  document.querySelectorAll(".ant").forEach((ant) => ant.remove());

  if (animationInterval) {
    console.log("Nettoyage de l'ancien interval");
    clearInterval(animationInterval);
  }
  
  console.log("Démarrage du setInterval avec vitesse:", speed);
  animationInterval = setInterval(nextTurn, 1000 / speed);
  console.log("Interval ID:", animationInterval);
}

let antFlipState = {};

function animateAntAlongPath(antElement, fromRoom, toRoom, duration, callback) {
  const emojiText = antElement.querySelector(".ant-emoji");
  const startX = fromRoom.x;
  const startY = fromRoom.y;
  const endX = toRoom.x;
  const endY = toRoom.y;
  
  const startTime = performance.now();
  
  function animate(currentTime) {
    const elapsed = currentTime - startTime;
    const progress = Math.min(elapsed / duration, 1);
    
    // Interpolation linéaire
    const currentX = startX + (endX - startX) * progress;
    const currentY = startY + (endY - startY) * progress;
    
    emojiText.setAttribute("x", currentX);
    emojiText.setAttribute("y", currentY);
    
    // Rotation basée sur la direction
    const angle = Math.atan2(endY - startY, endX - startX) * (180 / Math.PI);
    const antId = antElement.id.replace("ant-", "");
    
    // Dandiner pendant le mouvement
    const wiggleAngle = Math.sin(progress * Math.PI * 4) * 10;
    const rotation = angle + wiggleAngle;
    
    emojiText.setAttribute("transform", `rotate(${rotation} ${currentX} ${currentY})`);
    
    if (progress < 1) {
      requestAnimationFrame(animate);
    } else if (callback) {
      // Animation terminée, appeler le callback s'il existe
      callback();
    }
  }
  
  requestAnimationFrame(animate);
}

function pauseAntInRoom(antElement, room, pauseDuration) {
  // Effet de pulsation pendant la pause
  const emojiText = antElement.querySelector(".ant-emoji");
  const startTime = performance.now();
  const originalSize = 28;
  
  function pulse(currentTime) {
    const elapsed = currentTime - startTime;
    if (elapsed < pauseDuration) {
      const pulseProgress = (elapsed % 300) / 300;
      const scale = 1 + Math.sin(pulseProgress * Math.PI) * 0.2;
      const newSize = originalSize * scale;
      emojiText.setAttribute("font-size", newSize);
      requestAnimationFrame(pulse);
    } else {
      // Remettre la taille normale
      emojiText.setAttribute("font-size", originalSize);
    }
  }
  
  requestAnimationFrame(pulse);
}

function nextTurn() {
  if (isPaused || currentTurnIndex >= turns.length) {
    if (currentTurnIndex >= turns.length) {
      clearInterval(animationInterval);
    }
    return;
  }

  const turn = turns[currentTurnIndex];
  const moves = turn.split(" ");
  const svg = document.getElementById("graphSvg");
  
  console.log("Turn:", currentTurnIndex, "Moves:", moves);

  const activeAnts = new Set();
  const antsInEndRoom = new Set();

  moves.forEach((move) => {
    const [ant, room] = move.split("-");
    activeAnts.add(ant);
    
    console.log("Processing ant:", ant, "to room:", room);

    if (rooms[room]) {
      let antElement = document.getElementById(`ant-${ant}`);
      
      if (!antElement) {
        // Crée une nouvelle fourmi - elle part forcément de la salle de départ
        const g = document.createElementNS("http://www.w3.org/2000/svg", "g");
        g.setAttribute("class", "ant");
        g.setAttribute("id", `ant-${ant}`);
        
        // Trouve la salle de départ (type "start")
        let startRoom = null;
        for (let roomName in rooms) {
          if (rooms[roomName].type === "start") {
            startRoom = rooms[roomName];
            break;
          }
        }
        
        const emojiText = document.createElementNS("http://www.w3.org/2000/svg", "text");
        // Place la fourmi à la position de départ
        emojiText.setAttribute("x", startRoom.x);
        emojiText.setAttribute("y", startRoom.y);
        emojiText.setAttribute("font-size", "28");
        emojiText.setAttribute("text-anchor", "middle");
        emojiText.setAttribute("dominant-baseline", "central");
        emojiText.setAttribute("class", "ant-emoji");
        emojiText.textContent = "🐜";
        g.appendChild(emojiText);
        
        svg.appendChild(g);
        
        // Anime immédiatement vers la première destination
        const duration = 800 / speed;
        const isIntermediateRoom = rooms[room].type !== "start" && rooms[room].type !== "end";
        
        if (isIntermediateRoom && (room === "gilfoyle" || room === "dinesh" || room === "jimYoung")) {
          animateAntAlongPath(g, startRoom, rooms[room], duration, () => {
            pauseAntInRoom(g, rooms[room], 400 / speed);
          });
        } else {
          animateAntAlongPath(g, startRoom, rooms[room], duration);
        }
        
        antPreviousPositions[ant] = room;
      } else {
        // Anime le déplacement depuis la position précédente
        const previousRoom = antPreviousPositions[ant];
        if (previousRoom && rooms[previousRoom]) {
          const duration = 800 / speed; // Durée de l'animation en ms
          
          // Vérifier si la fourmi arrive à une salle intermédiaire (pas start ni end)
          const isIntermediateRoom = rooms[room].type !== "start" && rooms[room].type !== "end";
          
          if (isIntermediateRoom && (room === "gilfoyle" || room === "dinesh" || room === "jimYoung")) {
            // Animer jusqu'à la salle, puis faire une pause visible
            animateAntAlongPath(antElement, rooms[previousRoom], rooms[room], duration, () => {
              pauseAntInRoom(antElement, rooms[room], 400 / speed); // Pause de 400ms
            });
          } else {
            // Animation normale sans pause
            animateAntAlongPath(antElement, rooms[previousRoom], rooms[room], duration);
          }
        }
        antPreviousPositions[ant] = room;
      }

      // Marque les fourmis arrivées dans la salle de fin
      if (rooms[room].type === "end") {
        finishedAntsSet.add(ant);
      }
    }
  });

  // Supprime les fourmis qui ne sont plus actives (arrivées)
  document.querySelectorAll(".ant").forEach((antElement) => {
    const antId = antElement.id.replace("ant-", "");
    if (!activeAnts.has(antId)) {
      antElement.remove();
      delete antPreviousPositions[antId];
    }
  });

  // Met à jour les compteurs
  finishedAnts = finishedAntsSet.size;

  currentTurnIndex++;
  document.getElementById("currentTurn").textContent = currentTurnIndex;
  document.getElementById("activeAnts").textContent = activeAnts.size;
  document.getElementById("finishedAnts").textContent = finishedAnts;
}

function pauseAnimation() {
  isPaused = true;
}

function resumeAnimation() {
  isPaused = false;
}

function resetAnimation() {
  if (animationInterval) {
    clearInterval(animationInterval);
    animationInterval = null;
  }
  isPaused = false;
  currentTurnIndex = 0;
  finishedAnts = 0;
  finishedAntsSet.clear();
  antPreviousPositions = {};
  document.getElementById("currentTurn").textContent = "0";
  document.getElementById("activeAnts").textContent = "0";
  document.getElementById("finishedAnts").textContent = "0";

  const svg = document.getElementById("graphSvg");
  document.querySelectorAll(".ant").forEach((ant) => ant.remove());
}