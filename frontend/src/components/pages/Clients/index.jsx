import styles from './clients.module.css';
import { FaPlus } from "react-icons/fa";

function Clients() {
  return (
    <div>
      <h1 className={styles.title}>Clientes</h1>
      <section className={styles.container}>
        <input type="text" placeholder="Search" className={styles.container__input}/>
        <button type="button" className={styles.container__botton}><span><FaPlus /></span>Nuevo</button>
      </section>
      <div>
      <table>
          <thead>
            <tr>
              <th>Nombre</th>
              <th>Direccion</th>
              <th>Telefono</th>
              <th>Plan</th>
              <th>Fecha Pago</th>
              <th>Acciones</th>
            </tr>
          </thead>
          <tbody>
            <tr>
              <td>Stanley Hidalgo</td>
              <td>Villa Hermosa</td>
              <td>809-855-7908</td>
              <td>5M</td>
              <td>30</td>
              <td></td>
            </tr>
            <tr>
              <td>Stanley Hidalgo</td>
              <td>Villa Hermosa</td>
              <td>809-855-7908</td>
              <td>5M</td>
              <td>30</td>
              <td></td>
            </tr>
            <tr>
              <td>Stanley Hidalgo</td>
              <td>Villa Hermosa</td>
              <td>809-855-7908</td>
              <td>5M</td>
              <td>30</td>
              <td></td>
            </tr>
            <tr>
              <td>Stanley Hidalgo</td>
              <td>Villa Hermosa</td>
              <td>809-855-7908</td>
              <td>5M</td>
              <td>30</td>
              <td></td>
            </tr>
            <tr>
              <td>Stanley Hidalgo</td>
              <td>Villa Hermosa</td>
              <td>809-855-7908</td>
              <td>5M</td>
              <td>30</td>
              <td></td>
            </tr>
            <tr>
              <td>Stanley Hidalgo</td>
              <td>Villa Hermosa</td>
              <td>809-855-7908</td>
              <td>5M</td>
              <td>30</td>
              <td></td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  )
}

export default Clients