import requests
from pymongo import MongoClient

#从API接受json数据转化为字典
def url_data_2_dict(url):
    req = requests.get(url)
    req_dict = req.json()
    return req_dict
#预先询问mongodb地址端口
def prepare_func():
    host_addr = input("请输入host:")
    port = input("请输入端口号：")
    return host_addr, port

if __name__ == '__main__':
    #获得地址和端口
    host_addr, port = prepare_func()

    # 连接 MongoDB
    client = MongoClient(host=host_addr ,port=int(port))

    #通过对数据库列表长度进行判断是否连接成功
    if len(client.list_database_names()) == 0:
        print("连接失败")

    #连接到指定的数据库
    db = client.oyc_task

    #查找集合，如果有则写入数据，如果没有则创建集合
    collection = db.api_data
    #输入需要写入数据的API
    url = input("请输入url")

    #将API得到的数据写入指定数据库的指定集合中，返回_id,通过_id来判断是否写入成功。
    id = collection.insert_one(url_data_2_dict(url)).inserted_id
    if id != None :
        print("API数据写入MongoDB成功！")