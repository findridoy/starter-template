import { Table } from "antd";
import { useLoaderData } from "react-router-dom";

export default function UserIndex(){
  const {users} = useLoaderData()

  const columns = [
    {
      title: "Username",
      dataIndex: "username",
      key: "username"
    }
  ]
  return (
    <>
      <Table columns={columns} dataSource={users}/>
    </>
  )
}


