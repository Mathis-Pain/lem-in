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
    console.error("Erreur de chargement:", error);
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

  // Utilise les vrais liens du backend
  linksData.forEach((link) => {
    if (rooms[link.from] && rooms[link.to]) {
      paths.push({ from: link.from, to: link.to });
    }
  });

  // Compte les fourmis depuis les mouvements
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
  isPaused = false;
  document.getElementById("currentTurn").textContent = "0";
  document.getElementById("finishedAnts").textContent = "0";

  if (animationInterval) clearInterval(animationInterval);
  animationInterval = setInterval(nextTurn, 1000 / speed);
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
  document.querySelectorAll(".ant").forEach((ant) => ant.remove());

  const activeAnts = new Set();

  moves.forEach((move) => {
    const [ant, room] = move.split("-");
    activeAnts.add(ant);

    if (rooms[room]) {
      const g = document.createElementNS(
        "http://www.w3.org/2000/svg",
        "g"
      );
      g.setAttribute("class", "ant");

      const circle = document.createElementNS(
        "http://www.w3.org/2000/svg",
        "circle"
      );
      circle.setAttribute("cx", rooms[room].x);
      circle.setAttribute("cy", rooms[room].y);
      circle.setAttribute("r", 15);
      g.appendChild(circle);

      const text = document.createElementNS(
        "http://www.w3.org/2000/svg",
        "text"
      );
      text.setAttribute("x", rooms[room].x);
      text.setAttribute("y", rooms[room].y);
      text.textContent = ant.substring(1);
      g.appendChild(text);

      svg.appendChild(g);

      if (rooms[room].type === "end") {
        finishedAnts++;
      }
    }
  });

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
  document.getElementById("currentTurn").textContent = "0";
  document.getElementById("activeAnts").textContent = "0";
  document.getElementById("finishedAnts").textContent = "0";

  const svg = document.getElementById("graphSvg");
  document.querySelectorAll(".ant").forEach((ant) => ant.remove());
}