import React from 'react'
import ReactDOM from 'react-dom/client'
import { RouterProvider } from "react-router-dom"
import router from './router/Router.jsx'
import './fonts/Roboto/roboto.css'
import '../src/styles/index.css'
import '../src/styles/reset.css'

ReactDOM.createRoot(document.getElementById('root')).render(
  <React.StrictMode>
    <RouterProvider router={router} />
  </React.StrictMode>,
)
