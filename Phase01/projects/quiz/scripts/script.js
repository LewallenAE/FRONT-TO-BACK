// ==========================================
// 1. DATA
// ==========================================
const quizData = [
    { 
        question: "Is JavaScript a single-threaded language?", 
        answer: true, 
        explanation: "JavaScript is single-threaded because it has one call stack and one memory heap. This means exactly one command is executed at a time, and in sequential order!"
    },
    { 
        question: "Does `NaN === NaN` evaluate to true in JS?", 
        answer: false, 
        explanation: "This is the only value in JavaScript that is not equal to itself. Strictly defined by the IEEE 754, Not-a-Number is undefined, similar to dividing by zero, thus it is not representable as a mathematical result."
    }
];

// ==========================================
// 2. DOM ELEMENTS
// ==========================================
const questionBox = document.querySelector(".questionBox");
const answerButtons = document.querySelectorAll(".answerButton"); // Fixed: plural answerButtons
const nextButton = document.querySelector(".nextButton");
const scoreDisplay = document.querySelector(".score");

// ==========================================
// 3. STATE VARIABLES
// ==========================================
let currentIndex = 0;
let score = 0;
let hasAnswered = false;

// ==========================================
// 4. CORE FUNCTIONS
// ==========================================

function loadQuestion() {
    hasAnswered = false;
    resetButtons();

    if (currentIndex >= quizData.length) {
        showFinalResults();
        return;
    }

    const currentQuestion = quizData[currentIndex]; // Fixed: quizData instead of question
    questionBox.textContent = currentQuestion.question;
    scoreDisplay.textContent = `Score: ${score}/${quizData.length}`; // Fixed: quizData.length
    nextButton.textContent = "Next Question";
}

// Answer Button Clicks
answerButtons.forEach(button => {
    button.addEventListener('click', (event) => {
        if (hasAnswered) return;
        hasAnswered = true;

        const userChoice = event.target.getAttribute('data-value') === 'true';
        const currentQuestion = quizData[currentIndex]; // Fixed: quizData instead of question
        
        if (userChoice === currentQuestion.answer) {
            score++;
            event.target.style.background = 'var(--clr-seafoamGreen)';
            event.target.style.color = '#000';
        } else {
            event.target.style.background = 'var(--clr-deepWaterColorRed)'; // Fixed typo: event instead of even
            event.target.style.color = '#fff';

            answerButtons.forEach(btn => {
                if ((btn.getAttribute('data-value') === 'true') === currentQuestion.answer) {
                    btn.style.background = 'var(--clr-seafoamGreen)';
                    btn.style.color = '#000';
                }
            });
        }
        
        // Show current question + explanation and lock the buttons
        questionBox.textContent = `${currentQuestion.question}\n\n${currentQuestion.explanation}`;
        scoreDisplay.textContent = `Score: ${score}/${quizData.length}`;
        answerButtons.forEach(btn => btn.setAttribute('disabled', 'true'));
    });
});

// Reset styling and enable buttons for next question
function resetButtons() {
    answerButtons.forEach(btn => {
        btn.removeAttribute('style');
        btn.removeAttribute('disabled');
    });
}

// Next Button Click Listener
nextButton.addEventListener('click', () => {
    // Restart quiz if already complete
    if (currentIndex >= quizData.length) {
        currentIndex = 0;
        score = 0;
        answerButtons.forEach(btn => btn.style.display = 'grid');
        loadQuestion();
        return;
    }

    if (!hasAnswered) {
        alert("Please select an answer first!");
        return;
    }

    currentIndex++;
    loadQuestion();
});

// Show final results screen
function showFinalResults() {
    questionBox.textContent = `Quiz Complete! You scored ${score} out of ${quizData.length}.`;
    scoreDisplay.textContent = `Score: ${score}/${quizData.length}`;
    
    answerButtons.forEach(btn => btn.style.display = 'none');
    nextButton.textContent = "Restart Quiz";
}

// Initialize Question 1
loadQuestion();