import { Outlet } from "react-router-dom"
import styles from './app.module.css'
import Header from "../../organisms/Header/Index"
import SideBar from "../../organisms/SideBar"

function App() {
  return (
  <div className={styles.layout}>
    <div>
      <Header/>
    </div>
    <div className={styles.content}>
      <aside className={styles.content__sidebar}>
        <SideBar/>
      </aside>
      <main className={styles.content__main}>
        <Outlet/>
      </main>
    </div>
  </div>
  )
}

export default App