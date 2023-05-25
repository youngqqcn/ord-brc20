# brc20 checker
BRC20 token name availability checker written in like 5min


#### Instructions
- Enter token names in names.txt (example entries are already in there, each one should be on a new line)
- open a terminal/cmd in the folder with the checker
- Setup a virtual environment by using the command in terminal/cmd "python -m venv venv" (this will create a folder called "venv" in the same directory as the checker)
- activate the virtual environment by using the command in terminal/cmd "venv\Scripts\activate" (terminal should look slightly different after this, as you have activated the virtual env by using the activate script located inside that folder we just made)
- run the command in cmd/terminal "pip install -r requirements.txt" to get required packages
- run the command in cmd/terminal "python brc_checker.py" to run the checker after the required packages have been installed

- Hits/unclaimed names will be saved to the file "unclaimed.txt"
- On each run, the checker will add the new hits to the "unclaimed.txt" file if it already exists. If not, this file is created for you.

#### For support, please contact the discord: >_>#9249
