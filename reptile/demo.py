import requests
import re
import os
from bs4 import BeautifulSoup
import time


def gethtml(url, headers, proxies):   # 获取每日更新内容
    try:
        r = requests.get(url, headers=headers, proxies=proxies)
        r.raise_for_status()
        r.encoding = 'utf-8'
        return r.text
    except:
        print("请求失败")


def getnew(html):  # 获取最新15个内容返回对应的名字和地址
    #  soup = BeautifulSoup(html,'html.parser')
    #  for sibling in soup.p.next_sibling:
    #  a=sibling
    p = r'href="(.*?)"'
    name = r'target="_blank">(.*?)<'
    filenames = re.findall(name,html)
    urlm = re.findall(p, html)
    names = filenames[:15]  # 取前15个作为文件对象
    urls = urlm[15:30]  # 根据返回值确定15个地址
    return [urls, names]  # 返回一个数组


def get_img(newurls,names):  # 获得每个内容下的图片并保存在相映文件夹中
    headers = {
        "User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.102 Safari/537.36 Edge/18.18362",
        "Referer": ""}
    proxies = {'http': 'http://10.10.10.1:4321',
               'http': 'http:110.243.15.41:9999',
               'http': 'http:49.86.181.25:9999'}
    for i in range(0, 15):
        html = gethtml(newurls[i], headers, proxies)  # 获得每部分第一页的内容
        #  print(html)
        #  向请求头中添加防盗链信息，可以分析获得地址公式
        headers["Referer"] = newurls[i] + "/"
        try:

            for j in range(0, 80):  # 限制每个文件夹的图片数
                eveyurl = newurls[i] + "/" + str(j)
                #  print(eveyurl)
                r = requests.get(eveyurl, headers=headers, proxies=proxies)
                time.sleep(0.5)
                r.raise_for_status()
                r.encoding = r.apparent_encoding
                demo = r.text
                # 解析源码
                soup = BeautifulSoup(demo, "html.parser")
                # 得到外层标签
                div = soup.find("div", attrs={"class": "main-image"})
                # 得到图片地址
                img_url = div.find("img")["src"]
                print(img_url)

                # 请求图片地址，并获得响应
                photo = requests.get(img_url, headers=headers, proxies=proxies)
                time.sleep(0.5)
                # 保存图片到相应的文件夹
                with open("F:/lpthw/" + names[i] + "/" + str(j) + ".jpg", "wb") as f:
                    f.write(photo.content)
        except:
            continue  # 页数不够继续运行


def openfile(path, names):  # 根据名字创建文件夹
    """

    @param path:
    @param names:
    """
    for num in names:
        # 判断文件夹是否存在，如果存在不做处理，不存在则创建文件夹
        if os.path.exists(path + num):
            pass
        else:
            os.mkdir(path + num)


def main():
    url = "https://www.mzitu.com/all/"
    headers = {"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.102 Safari/537.36 Edge/18.18362",
              "Referer": "https://www.mzitu.com/"}
    proxies = {'http': 'http://10.10.10.1:4321',
               'http': 'http:110.243.15.41:9999',
               'http': 'http:49.86.181.25:9999'}
    html = gethtml(url, headers, proxies)  # 获得每日最新页面内容
    newurls = getnew(html)[0]
    names = getnew(html)[1]
    raw_path = "F:/lpthw/"
    openfile(raw_path, names)
    get_img(newurls, names)

    print(newurls)
    #  print(names)


if __name__ == '__main__':
    # demo 这个网址已经被和谐了，
    main()
