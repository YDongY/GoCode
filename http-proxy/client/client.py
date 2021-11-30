import requests


def main(url):
    proxies = {
        "http": "http://localhost:8080"
    }

    headers = {
        "Name": "Mark",
        "Via": "HTTP/1.1 Test, HTTP/1.1 Test2"
    }

    resp = requests.get(url, headers=headers, proxies=proxies)
    assert resp.status_code == 200


if __name__ == '__main__':
    main("http://www.baidu.com")
