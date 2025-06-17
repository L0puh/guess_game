![](data/circuit.png)

# IDEA
The system generates randomized light sequences that players must observe and reproduce accurately by inputting the correct order.
# STACK
- arduino (3 LEDs, 3 buttons)
- Golang (net/http) for web interface
- Python for utils 
# RUN
- connect arduino (make sure it's running on `tty/USB0`
- run the script: `./run.sh`
- open `localhost:8080`
