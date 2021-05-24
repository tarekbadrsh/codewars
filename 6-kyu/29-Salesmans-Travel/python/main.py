

def travel(r, zipcode):
    str_info = []
    str_numbers = [] 
    for str in r.split(","):
        tmp = str.strip().split(" ")
        if zipcode == " ".join(tmp[(len(tmp)-2):]): 
            str_numbers.append(tmp[0])
            str_info.append(" ".join(tmp[1:len(tmp)-2])) 
    return zipcode + ":" + ",".join(str_info) + "/" + ",".join(str_numbers)

print(travel("123 Main Street St. Louisville OH 43071,432 Main Long Road St. Louisville OH 43071,786 High Street Pollocksville NY 56432","OH 43071"))
