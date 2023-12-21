import styles from './service.module.css';

function Service() {
  return (
    <div>
      <h1 className={styles.service__title}>Service</h1>
      <div className={styles.container}>
        <section className={styles.service__table_container}>
          <table>
            <thead>
              <tr>
                <th>Nombre</th>
                <th>Velovidad</th>
                <th>Precio</th>
                <th>Acciones</th>
              </tr>
            </thead>
            <tbody>
              <tr>
                <td>5M</td>
                <td>5000/1500 <span>Kbs</span></td>
                <td><span>$</span>800</td>
                <td></td>
              </tr>
              <tr>
                <td>5M</td>
                <td>5000/1500 <span>Kbs</span></td>
                <td><span>$</span>800</td>
                <td></td>
              </tr>
              <tr>
                <td>5M</td>
                <td>5000/1500 <span>Kbs</span></td>
                <td><span>$</span>800</td>
                <td></td>
              </tr>
              <tr>
                <td>5M</td>
                <td>5000/1500 <span>Kbs</span></td>
                <td><span>$</span>800</td>
                <td></td>
              </tr>
            </tbody>
          </table>
        </section>
        
        <section className={styles.service__form_container}>
          <form className={styles.service__form}>
            
            <h2 className={styles.form__title}>Nuevo Servicio</h2>

            <label htmlFor="name">Nombre</label>
            <input type="text" name="name"/>

            <label htmlFor="velocidad">Velocidad</label>
            <input type="text" name="velovidad"/>
            
            <label htmlFor="precio">Precio</label>
            <input type="text" name="precio"/>

            <div className={styles.form__buttons}>
              <button type="reset" className={`${styles.form__button} ${styles.color__red}`}>Cancelar</button>
              <button type="submit" className={`${styles.form__button} ${styles.color__primary}`}>Guardar</button>
            </div>
          </form>
        </section>
      </div>
    </div>
  )
}

export default Service