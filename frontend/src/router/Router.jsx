import { createBrowserRouter } from "react-router-dom"
import App from "../components/templates/App/index"
import Error404 from "../components/pages/Error404"
import Dashboard from "../components/pages/Dashboard/index"
import Clients from "../components/pages/Clients/index"
import Bills from "../components/pages/Bills/index"
import Service from "../components/pages/Service/index"
import Reports from "../components/pages/Reports/index"
const router = createBrowserRouter([
    {
        path:"/",
        element:<App/>,
        errorElement:<Error404/>,
        children: [
            {
                index:true,
                element:<Dashboard/>,
            },
            {
                path:"/clients",
                element:<Clients/>,
            },
            {
                path:"/bills",
                element:<Bills/>,
            },
            {
                path:"/services",
                element:<Service/>,
            },
            {
                path:"/reports",
                element:<Reports/>,
            },
        ]
    }
])


export default router
