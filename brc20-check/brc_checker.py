import requests, threading, time, os
from bs4 import BeautifulSoup


def check(token):
    while True:
        endpoint = f"https://brc-20.io/token?n={token}"
        with requests.Session() as s:
            try:
                h = {"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:109.0) Gecko/20100101 Firefox/112.0", "Accept": "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8", "Accept-Language": "en-US,en;q=0.5",  "Referer": "https://brc-20.io/", "DNT": "1", "Connection": "keep-alive", "Upgrade-Insecure-Requests": "1", "Sec-Fetch-Dest": "document", "Sec-Fetch-Mode": "navigate", "Sec-Fetch-Site": "same-origin", "Sec-Fetch-User": "?1", "Cache-Control": "max-age=0", "TE": "trailers"}
                r = s.get(endpoint, headers=h)
                soup = BeautifulSoup(r.text, "lxml")
                #print(soup.prettify())
                #time.sleep(15)
                # if r.status_code == 200: and the token name is in the html soup then it is registered
                if r.status_code == 200 and token in soup.text:
                    print("\u001b[31m" + token + " is claimed.\n")
                    break
                else:
                    # if the token name is not in the html soup then it is not registered
                    if token not in soup.text:
                        print("\u001b[32m" + token + " is not claimed.\n")
                        with open("unclaimed.txt", "a+") as f:
                            f.write(token + "\n")
                        break
            except Exception as e:
                print("\u001b[31mencountered error checking token - " + str(token) + "\n")
                print(e)
                break

def loader(threadcount):
    with open("names.txt", "r") as f:
        names = [x.split() for x in f.read().splitlines()]
    print("\nStarting...Loaded " + str(len(names)) + " tokens to check.")
    for n in names:
        while True:
            if threading.active_count() < threadcount:
                threading.Thread(target=check, args=(n)).start()
                break
            time.sleep(1)
        
os.system('cls' if os.name == 'nt' else 'clear')
art = """\u001b[33m
 _                     _               _             
| |__  _ __ ___    ___| |__   ___  ___| | _____ _ __ 
| '_ \| '__/ __|  / __| '_ \ / _ \/ __| |/ / _ \ '__|
| |_) | | | (__  | (__| | | |  __/ (__|   <  __/ |   
|_.__/|_|  \___|  \___|_| |_|\___|\___|_|\_\___|_|   
                                                     
"""
print(art)
ask = int(input('\u001b[33m-Enter 1 to start\n-Or 2 to exit.\n\n'))
threadcount = int(input('\u001b[33m-Enter the amount of threads to use\n(recommended is 3-5. do not set this too high to avoid possible ratelimits).\n\n'))
if ask == 1:
    print('Starting...')
    time.sleep(0.5)
    print()
    print()
    loader(threadcount)
elif ask == 2:
    exit()
else:
    print('invalid input, closing program.')
    exit()
        