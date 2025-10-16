let turns = [];
let currentTurnIndex = 0;
let animationInterval = null;
let isPaused = false;
let speed = 1;
let rooms = {};
let paths = [];
let totalAnts = 0;
let finishedAnts = 0;
let finishedAntsSet = new Set();
let antPreviousPositions = {};

window.onload = async function () {
  try {
    const roomsResponse = await fetch("/api/rooms");
    const roomsData = await roomsResponse.json();

    const movesResponse = await fetch("/api/moves");
    const movesDataResponse = await movesResponse.json();

    const linksResponse = await fetch("/api/links");
    const linksData = await linksResponse.json();

    if (movesDataResponse.moves) {
      turns = movesDataResponse.moves
        .split("\n")
        .filter((line) => line.trim());

      if (roomsData && Object.keys(roomsData).length > 0) {
        loadRoomsWithCoordinates(roomsData, linksData);
      }

      createGraph();
      document.querySelector(".loading").style.display = "none";
      document.getElementById("totalTurns").textContent = turns.length;
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

  // Ajoute des dégradés pour les trous
  const defs = document.createElementNS("http://www.w3.org/2000/svg", "defs");
  
  // Dégradé vert pour le départ
  const startGradient = document.createElementNS("http://www.w3.org/2000/svg", "radialGradient");
  startGradient.setAttribute("id", "startGradient");
  const startStop1 = document.createElementNS("http://www.w3.org/2000/svg", "stop");
  startStop1.setAttribute("offset", "0%");
  startStop1.setAttribute("style", "stop-color:#1a3d1a;stop-opacity:1");
  const startStop2 = document.createElementNS("http://www.w3.org/2000/svg", "stop");
  startStop2.setAttribute("offset", "100%");
  startStop2.setAttribute("style", "stop-color:#228B22;stop-opacity:1");
  startGradient.appendChild(startStop1);
  startGradient.appendChild(startStop2);
  
  // Dégradé rouge pour la fin
  const endGradient = document.createElementNS("http://www.w3.org/2000/svg", "radialGradient");
  endGradient.setAttribute("id", "endGradient");
  const endStop1 = document.createElementNS("http://www.w3.org/2000/svg", "stop");
  endStop1.setAttribute("offset", "0%");
  endStop1.setAttribute("style", "stop-color:#4d0000;stop-opacity:1");
  const endStop2 = document.createElementNS("http://www.w3.org/2000/svg", "stop");
  endStop2.setAttribute("offset", "100%");
  endStop2.setAttribute("style", "stop-color:#8B0000;stop-opacity:1");
  endGradient.appendChild(endStop1);
  endGradient.appendChild(endStop2);
  
  defs.appendChild(startGradient);
  defs.appendChild(endGradient);
  svg.appendChild(defs);

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
    clearInterval(animationInterval);
  }
  
  animationInterval = setInterval(nextTurn, 1000 / speed);
}

function animateAntAlongPath(antElement, fromRoom, toRoom, fromRoomName, toRoomName, duration, callback) {
  const antSvg = antElement.querySelector(".ant-svg");
  const startX = fromRoom.x;
  const startY = fromRoom.y;
  const endX = toRoom.x;
  const endY = toRoom.y;
  
  const startTime = performance.now();
  
  // Calcule l'angle de la ligne droite entre les deux cercles
  const lineAngle = Math.atan2(endY - startY, endX - startX) * (180 / Math.PI);
  
  function animate(currentTime) {
    const elapsed = currentTime - startTime;
    const progress = Math.min(elapsed / duration, 1);
    
    // Interpolation linéaire
    const currentX = startX + (endX - startX) * progress;
    const currentY = startY + (endY - startY) * progress;
    
    // Petit dandinage
    const wiggleAngle = Math.sin(progress * Math.PI * 4) * 5;
    
    // Rotation : angle de la ligne + 90° (fourmi orientée vers le haut par défaut)
    const finalRotation = lineAngle + 90 + wiggleAngle;
    
    // Applique la transformation : translate puis rotate
    const transform = `translate(${currentX}, ${currentY}) rotate(${finalRotation})`;
    antSvg.setAttribute("transform", transform);
    
    if (progress < 1) {
      requestAnimationFrame(animate);
    } else if (callback) {
      callback();
    }
  }
  
  requestAnimationFrame(animate);
}

function nextTurn() {
  if (isPaused || currentTurnIndex >= turns.length) {
    if (currentTurnIndex >= turns.length) {
      clearInterval(animationInterval);
      // Attends que toutes les animations se terminent (800ms + délai)
      setTimeout(() => {
        document.getElementById("activeAnts").textContent = "0";
        document.getElementById("finishedAnts").textContent = finishedAntsSet.size;
      }, 1000 / speed);
    }
    return;
  }

  const turn = turns[currentTurnIndex];
  const moves = turn.split(" ");
  const svg = document.getElementById("graphSvg");

  const activeAnts = new Set();

  moves.forEach((move) => {
    const [ant, room] = move.split("-");
    activeAnts.add(ant);
    
    if (rooms[room]) {
      let antElement = document.getElementById(`ant-${ant}`);
      
      if (!antElement) {
        const g = document.createElementNS("http://www.w3.org/2000/svg", "g");
        g.setAttribute("class", "ant");
        g.setAttribute("id", `ant-${ant}`);
        
        let startRoom = null;
        let startRoomName = null;
        for (let roomName in rooms) {
          if (rooms[roomName].type === "start") {
            startRoom = rooms[roomName];
            startRoomName = roomName;
            break;
          }
        }
        
        // Crée une fourmi SVG
        const antSvg = document.createElementNS("http://www.w3.org/2000/svg", "g");
        antSvg.setAttribute("class", "ant-svg");
        
        // Corps de la fourmi (3 ellipses) - orientée vers le haut par défaut
        const head = document.createElementNS("http://www.w3.org/2000/svg", "ellipse");
        head.setAttribute("cx", "0");
        head.setAttribute("cy", "-8");
        head.setAttribute("rx", "3");
        head.setAttribute("ry", "4");
        head.setAttribute("fill", "#2c1810");
        
        const thorax = document.createElementNS("http://www.w3.org/2000/svg", "ellipse");
        thorax.setAttribute("cx", "0");
        thorax.setAttribute("cy", "0");
        thorax.setAttribute("rx", "4");
        thorax.setAttribute("ry", "5");
        thorax.setAttribute("fill", "#3d2415");
        
        const abdomen = document.createElementNS("http://www.w3.org/2000/svg", "ellipse");
        abdomen.setAttribute("cx", "0");
        abdomen.setAttribute("cy", "8");
        abdomen.setAttribute("rx", "5");
        abdomen.setAttribute("ry", "7");
        abdomen.setAttribute("fill", "#2c1810");
        
        // Pattes
        for (let i = 0; i < 6; i++) {
          const leg = document.createElementNS("http://www.w3.org/2000/svg", "line");
          const side = i % 2 === 0 ? -1 : 1;
          const yPos = -3 + Math.floor(i / 2) * 3;
          leg.setAttribute("x1", "0");
          leg.setAttribute("y1", yPos);
          leg.setAttribute("x2", side * 8);
          leg.setAttribute("y2", yPos + 5);
          leg.setAttribute("stroke", "#2c1810");
          leg.setAttribute("stroke-width", "1.5");
          antSvg.appendChild(leg);
        }
        
        // Antennes
        const antenna1 = document.createElementNS("http://www.w3.org/2000/svg", "line");
        antenna1.setAttribute("x1", "0");
        antenna1.setAttribute("y1", "-12");
        antenna1.setAttribute("x2", "-3");
        antenna1.setAttribute("y2", "-16");
        antenna1.setAttribute("stroke", "#2c1810");
        antenna1.setAttribute("stroke-width", "1");
        
        const antenna2 = document.createElementNS("http://www.w3.org/2000/svg", "line");
        antenna2.setAttribute("x1", "0");
        antenna2.setAttribute("y1", "-12");
        antenna2.setAttribute("x2", "3");
        antenna2.setAttribute("y2", "-16");
        antenna2.setAttribute("stroke", "#2c1810");
        antenna2.setAttribute("stroke-width", "1");
        
        // Assemble la fourmi
        antSvg.appendChild(abdomen);
        antSvg.appendChild(thorax);
        antSvg.appendChild(head);
        antSvg.appendChild(antenna1);
        antSvg.appendChild(antenna2);
        
        // Positionne la fourmi au départ
        antSvg.setAttribute("transform", `translate(${startRoom.x}, ${startRoom.y})`);
        
        g.appendChild(antSvg);
        
        svg.appendChild(g);
        
        const duration = 800 / speed;
        
        // Si c'est la destination finale, ajoute un callback pour supprimer après
        if (rooms[room].type === "end") {
          animateAntAlongPath(g, startRoom, rooms[room], startRoomName, room, duration, () => {
            finishedAntsSet.add(ant);
            // Met à jour le compteur immédiatement quand la fourmi arrive
            document.getElementById("finishedAnts").textContent = finishedAntsSet.size;
            
            // Supprime après 0.5 seconde
            setTimeout(() => {
              activeAnts.delete(ant);
              const antToRemove = document.getElementById(`ant-${ant}`);
              if (antToRemove) {
                antToRemove.remove();
                delete antPreviousPositions[ant];
              }
              const activeCount = activeAnts.size - finishedAntsSet.size;
              document.getElementById("activeAnts").textContent = Math.max(0, activeCount);
            }, 500 / speed);
          });
        } else {
          animateAntAlongPath(g, startRoom, rooms[room], startRoomName, room, duration);
        }
        
        antPreviousPositions[ant] = room;
      } else {
        const previousRoom = antPreviousPositions[ant];
        if (previousRoom && rooms[previousRoom]) {
          const duration = 800 / speed;
          
          // Si c'est la destination finale, ajoute un callback pour supprimer après
          if (rooms[room].type === "end") {
            animateAntAlongPath(antElement, rooms[previousRoom], rooms[room], previousRoom, room, duration, () => {
              finishedAntsSet.add(ant);
              // Met à jour le compteur immédiatement quand la fourmi arrive
              document.getElementById("finishedAnts").textContent = finishedAntsSet.size;
              
              // Supprime après 0.5 seconde
              setTimeout(() => {
                activeAnts.delete(ant);
                const antToRemove = document.getElementById(`ant-${ant}`);
                if (antToRemove) {
                  antToRemove.remove();
                  delete antPreviousPositions[ant];
                }
                const activeCount = activeAnts.size - finishedAntsSet.size;
                document.getElementById("activeAnts").textContent = Math.max(0, activeCount);
              }, 500 / speed);
            });
          } else {
            animateAntAlongPath(antElement, rooms[previousRoom], rooms[room], previousRoom, room, duration);
          }
        }
        antPreviousPositions[ant] = room;
      }
    }
  });

  // Supprime les fourmis qui ne sont plus actives (sauf celles qui vont vers end ou sont dans end)
  document.querySelectorAll(".ant").forEach((antElement) => {
    const antId = antElement.id.replace("ant-", "");
    // Ne pas supprimer si la fourmi est finie (elle a son propre timer de suppression)
    if (!activeAnts.has(antId) && !finishedAntsSet.has(antId)) {
      antElement.remove();
      delete antPreviousPositions[antId];
    }
  });

  // Met à jour les compteurs
  // Pour le compteur d'actifs, on soustrait les fourmis finies
  const activeCount = activeAnts.size - finishedAntsSet.size;
  finishedAnts = finishedAntsSet.size;

  currentTurnIndex++;
  document.getElementById("currentTurn").textContent = currentTurnIndex;
  document.getElementById("activeAnts").textContent = Math.max(0, activeCount);
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