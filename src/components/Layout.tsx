import { Layout, Menu } from "antd";
import { Form, Outlet, redirect, useNavigate } from "react-router-dom";
import { Button } from "./ui/button";
import { ofetch } from "@/lib/ofetch";
import { useState } from "react";

export default function MainLayout() {
  const navigate = useNavigate();
  const { Header, Content, Footer, Sider } = Layout;

  const [menuItems, setMenuItems] = useState([
    {
      label: "Dashboard",
      key: "/dashboard",
    },
    {
      label: "Users",
      key: "/users",
    }
  ])


  const handleMenuClick = ({key}) => {
    // console.log(key);
    navigate(key)
  }


  return (
    <>
      <Layout hasSider>
        <Sider
          style={{ overflow: 'auto', height: '100vh', position: 'fixed', left: 0, top: 0, bottom: 0 }}
          theme="light"
        >
          <div className="relative h-full w-full">
            <Menu items={menuItems} onClick={handleMenuClick}/>

            <div className="w-full">
              <Form method="POST">
                <Button className="absolute bottom-0 left-0 right-0 mx-6 mb-8" type="submit">Logout</Button>
              </Form>
            </div>

          </div>
        </Sider>
        <Layout style={{ marginLeft: 200 }} className="bg-white">
          <Content style={{ margin: '4px 8px' }}>
            <Outlet />
          </Content>
        </Layout>
      </Layout>
    </>
  )
}


export const LayoutRouteAction = async ({ request }) => {
  console.log("yes");
  
  try {
    await ofetch("/logout", {method: "post"})
    document.cookie = "Authorization="
    return redirect("/login")
    
  } catch (error) {
    
  }

  return null
}
