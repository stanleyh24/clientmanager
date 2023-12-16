import { createBrowserRouter } from "react-router-dom"
import App from "../components/templates/App/index"
import Error404 from "../components/pages/Error404"
import Dashboard from "../components/pages/Dashboard"
import Clients from "../components/pages/Clients"
import Bills from "../components/pages/Bills"
import Service from "../components/pages/Service"
import Reports from "../components/pages/Reports"
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
