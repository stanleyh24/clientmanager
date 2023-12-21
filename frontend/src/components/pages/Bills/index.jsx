import styles from './bills.module.css';

function Bills() {
  return (
    <div className={styles.container__bills}>
      <h1 className={styles.bills__title}>Facturas</h1>

      <section className={styles.bills__search}>
        <input type="text" placeholder="Search" className={styles.search__input}/>
      </section>

      <section>
      <table>
          <thead>
            <tr>
              <th>Nombre</th>
              <th>Direccion</th>
              <th>Monto</th>
              <th>Plan</th>
              <th>Fecha Corte</th>
              <th>Acciones</th>
            </tr>
          </thead>
          <tbody>
            <tr>
              <td>Stanley Hidalgo</td>
              <td>Villa Hermosa</td>
              <td><span>$</span>800</td>
              <td>5M</td>
              <td>21-09-2023</td>
              <td></td>
            </tr>
            <tr>
              <td>Stanley Hidalgo</td>
              <td>Villa Hermosa</td>
              <td><span>$</span>800</td>
              <td>5M</td>
              <td>21-09-2023</td>
              <td></td>
            </tr>
            <tr>
              <td>Stanley Hidalgo</td>
              <td>Villa Hermosa</td>
              <td><span>$</span>800</td>
              <td>5M</td>
              <td>21-09-2023</td>
              <td></td>
            </tr>
            <tr>
              <td>Stanley Hidalgo</td>
              <td>Villa Hermosa</td>
              <td><span>$</span>800</td>
              <td>5M</td>
              <td>21-09-2023</td>
              <td></td>
            </tr>
            <tr>
              <td>Stanley Hidalgo</td>
              <td>Villa Hermosa</td>
              <td><span>$</span>800</td>
              <td>5M</td>
              <td>21-09-2023</td>
              <td></td>
            </tr>
          </tbody>
        </table>
      </section>
    </div>
  )
}

export default Bills