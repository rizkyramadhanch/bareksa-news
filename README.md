# bareksa-news
simple crud using gin golang

API Documentation

## Production Mode

### Home Page
[GET] `https://bareksa.herokuapp.com/`
Response : 
{
"Status": "OK"
}

### Handling No Route

[GET] `https://bareksa.herokuapp.com/no-route`

Response : 
{
"message": "Page not found"
}

### Get All News

[GET] `https://bareksa.herokuapp.com/news/list`

Response : 
{
    "data": [
        {
            "id": 3,
            "title": "How to start investment",
            "description": "For many people, the word “investing” conjures up images of men in suits, monitoring the exchange of millions of dollars on a stock ticker.\nI’m here to tell you: You don’t need to be the Wolf of Wall Street to start investing. It’s okay if you’re more of a mouse of Main Street. Even if you only have a few dollars to spare, your money will grow with compound interest.\n\nThe key to building wealth is developing good habits—like regularly putting money away every month. Swap out the barista-made cappuccinos for coffee at home and you could already be saving more than $50 a month.\n\nOnce you have a little money to play with, you can start to invest.\n\nIn 2020, you can get a date, a ride or a pizza with the swipe of a smartphone screen. Investing is no different. If you can automate your bills, why not your investments? It’s just as easy.\n\nWith a robo-advisor, you can make your MONEY work while you play. And just like Halloween costumes, investing comes in many different forms. It shouldn’t be a scary word.\n\nWhether it’s opening a savings account, investing in your retirement or the real estate market, investing for beginners is simpler and more straightforward than ever before.\n\nSoon you’ll see how addictive growing your money can be.",
            "status": "publish",
            "topic": "investment",
            "tag": null
        },
        {
            "id": 5,
            "title": "test",
            "description": "ajdajdajdjadkjakdjkajdad",
            "status": "publish",
            "topic": "investment",
            "tag": [
                "Investment"
            ]
        },
        {
            "id": 13,
            "title": "FInal Test",
            "description": "loremsjaljsfajfjasfjafja;",
            "status": "publish",
            "topic": "investment",
            "tag": [
                "mutual fund updated"
            ]
        },
        {
            "id": 1,
            "title": "What is Investment Banking?",
            "description": "Investment banking is the division of a bank or financial institution that serves governments, corporations, and institutions by providing underwriting (capital raising) and mergers and acquisitions (M&A) advisory services. Investment banks act as intermediaries between investors (who have money to invest) and corporations (who require capital to grow and run their businesses). This guide will cover what investment banking is and what investment bankers actually do.",
            "status": "publish",
            "topic": "investment",
            "tag": [
                "Investment",
                "mutual fund updated",
                "Ini Tag Baru"
            ]
        },
        {
            "id": 15,
            "title": "COba tambah",
            "description": "sisinya",
            "status": "publish",
            "topic": "market",
            "tag": null
        },
        {
            "id": 14,
            "title": "hahahha",
            "description": "hihihihh",
            "status": "publish",
            "topic": "market",
            "tag": [
                "Investment",
                "mutual fund updated"
            ]
        },
        {
            "id": 12,
            "title": "Investasi Emas",
            "description": "lorem ipsum bla blabla",
            "status": "publish",
            "topic": "market",
            "tag": null
        }
    ],
    "message": "Getting list of news successfully"
}

Get News by ID 

require newsID on url

### [GET] `https://bareksa.herokuapp.com/news/detail/{newsID}`

Response : 
{
    "data": {
        "id": 1,
        "title": "What is Investment Banking?",
        "description": "Investment banking is the division of a bank or financial institution that serves governments, corporations, and institutions by providing underwriting (capital raising) and mergers and acquisitions (M&A) advisory services. Investment banks act as intermediaries between investors (who have money to invest) and corporations (who require capital to grow and run their businesses). This guide will cover what investment banking is and what investment bankers actually do.",
        "status": "publish",
        "topic": "investment",
        "tag": [
            "Investment",
            "mutual fund updated",
            "Ini Tag Baru"
        ]
    },
    "message": "Getting news detail succesfully"
}

Update News

require newsID on url

### [POST] `https://bareksa.herokuapp.com/news/update/{newsID}`

require JSON body : 
{
    "title": "New Title",
    "description": "New description",
    "status_id": {statusID},
    "topic_id": 1
}

Response : 
{
    "data": "News ID 1 successfully updated",
    "message": "Update news successfully"
}

Add News 

require newsID on url

### [POST] `https://bareksa.herokuapp.com/news/add`

require JSON body : 
{
	"title"	: "New title",
	"description" : "New Description",
	"status_id" :1,
	"topic_id" :1
}

Response : 
{
    "data": "A news with title Coba tambah lagihas been created",
    "message": "Create a news succesfully"
}




