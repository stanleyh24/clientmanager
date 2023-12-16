import { Link } from "react-router-dom";
import styles from "./sidebar.module.css";
function SideBar() {
  return (
    <nav className={styles.menu}>
        <ul className={styles.menu__list}>
            <li className={styles.menu__option}>
                <Link to='/' className={styles.menu__link}>Dashboard</Link>
            </li>
        
            <li className={styles.menu__option}>
                <Link to='/clients' className={styles.menu__link}>Clientes</Link>
            </li>
        
            <li className={styles.menu__option}>
                <Link to='/bills' className={styles.menu__link}>Facturas</Link>
            </li>
        
            <li className={styles.menu__option}>
                <Link to='/services' className={styles.menu__link}>Servicios</Link>
            </li>
        
            <li className={styles.menu__option}>
                <Link to='/reports' className={styles.menu__link}>Reportes</Link>
            </li>
        </ul>
    </nav>
  )
}

export default SideBar