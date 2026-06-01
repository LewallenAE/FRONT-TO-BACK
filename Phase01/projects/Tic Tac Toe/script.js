const squares = document.querySelectorAll(".square");
let players = document.querySelectorAll(".player1, .player2");
let boardState = Array(9).fill(null);
let activeGame = true;
console.log("JS sees", players);

const player1 = {
    name: players[0].textContent,
    marker: 'X'
};

const player2 = {
    name: players[1].textContent,
    marker: 'O'
};

let currentPlayer = player1;

function checkWinner() {
    const winningCombinations = [
        [0, 1, 2], [3, 4, 5], [6, 7, 8], // Row Winners
        [0, 3, 6], [1, 4, 7], [2, 5, 8], // Column Winners
        [0, 4, 8], [2, 4, 6] // Diagonal Winners
    ];

    for (const combinations of winningCombinations) {

        const [a, b, c] = combinations;

        if (boardState[a] && boardState[a] === boardState[b] && boardState[a] === boardState[c]) { 
            return true;        
        }
    }
    return false;
};

squares.forEach((square, index) => {
    square.addEventListener('click', () => {
        
        if (!activeGame || boardState[index] !== null) return;

        boardState[index] = currentPlayer.marker;
        square.textContent = currentPlayer.marker;

        if (checkWinner()) {
            console.log(`${currentPlayer.name} wins!!!`);
            activeGame = false;
            return;
        } 

        currentPlayer = currentPlayer === player1 ? player2 : player1;
    });
});