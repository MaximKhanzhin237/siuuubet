document.addEventListener('DOMContentLoaded', function () {
    const Balance='{{.Balance}}';
    const betForm = document.getElementById('betForm');
    const betAmountInput = document.getElementById('betAmount');
    const betChoice = document.getElementById('betChoice');
    const potentialWinningsDisplay = document.getElementById('potentialWinnings');

    // Function to show league events
    window.showLeague = function (league) {
        const premierLeagueSection = document.getElementById('premier-league');
        const otherLeaguesSection = document.getElementById('other-leagues');

        if (league === 'premier') {
            premierLeagueSection.classList.remove('d-none');
            otherLeaguesSection.classList.add('d-none');
        } else {
            premierLeagueSection.classList.add('d-none');
            otherLeaguesSection.classList.remove('d-none');
        }
    };

    // Calculate potential winnings
    betAmountInput.addEventListener('input', function () {
        const odds = Array.from(betChoice.selectedOptions).map(option => parseFloat(option.getAttribute('data-odds')));
        const betAmount = parseFloat(betAmountInput.value);

        if (!isNaN(betAmount)) {
            const potentialWinnings = odds.map(odd => (betAmount * odd).toFixed(2));
            potentialWinningsDisplay.textContent = potentialWinnings.join(', ');
        } else {
            potentialWinningsDisplay.textContent = '0';
        }
    });

    // Handle bet submission
    betForm.addEventListener('submit', function () {
        const selectedOption = betChoice.options[betChoice.selectedIndex];
        const outcome = selectedOption.text;
        const odds = parseFloat(selectedOption.getAttribute('data-odds'));
        const betAmount = parseFloat(betAmountInput.value);
        const balance = Balance;

        if (balance >= betAmount && odds > 0) {
            const bet={"result": outcome, "bet_sum": betAmount, "odds": odds, "balance": balance};
            const Http = new XMLHttpRequest();
            const url='https://localhost:8080';
            Http.open("POST", url);
            Http.setRequestHeader("Content-Type", "application/json");
            Http.send(JSON.stringify(bet));
        } else {
            alert(balance);
        }
    });

    // Open bet modal with correct odds
    document.querySelectorAll('.bet-btn').forEach(button => {
        button.addEventListener('click', function () {
            const oddsElements = this.parentElement.querySelectorAll('.odds');
            const oddsArray = Array.from(oddsElements).map(odd => {
                return { outcome: odd.parentElement.textContent, odds: odd.dataset.odds };
            });

            // Clear previous options
            betChoice.innerHTML = '';

            // Add options with correct odds
            oddsArray.forEach(odd => {
                const option = document.createElement('option');
                option.value = odd.outcome;
                option.setAttribute('data-odds', odd.odds);
                option.textContent = `${odd.outcome} (Odds: ${odd.odds})`;
                betChoice.appendChild(option);
            });
        });
    });

    // Show the initial league
    showLeague('premier');
});


