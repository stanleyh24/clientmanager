import Cards from "../../organisms/Cards";
import styles from './dashboard.module.css';
function Dashboard() {
  return (
    <>
      <Cards/>

      <section className={styles.paid__container}>
        <h1>Ultimos Pagos</h1>
        <div className={styles.table__container}>
          <table >
            <thead>
              <tr>
                <th>Nombre</th>
                <th>Fecha</th>
                <th>Monto</th>
              </tr>
            </thead>
            <tbody>
              <tr>
                <td>Stanley Hidalgo</td>
                <td>15-09-2023</td>
                <td><span>$</span>800</td>
              </tr>
              <tr>
                <td>Stanley Hidalgo</td>
                <td>15-09-2023</td>
                <td><span>$</span>800</td>
              </tr>
              <tr>
                <td>Stanley Hidalgo</td>
                <td>15-09-2023</td>
                <td><span>$</span>800</td>
              </tr>
              <tr>
                <td>Stanley Hidalgo</td>
                <td>15-09-2023</td>
                <td><span>$</span>800</td>
              </tr>
              <tr>
                <td>Stanley Hidalgo</td>
                <td>15-09-2023</td>
                <td><span>$</span>800</td>
              </tr>
            </tbody>
          </table>
        </div>
      </section>
      
    </>
  )
}

export default Dashboard