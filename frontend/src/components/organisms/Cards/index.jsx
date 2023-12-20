import { FaUser } from "react-icons/fa";
import { BsCashStack } from "react-icons/bs";
import { FaUserTimes } from "react-icons/fa";
import styles from './cards.module.css';

function Cards() {
  return (
    <section className={styles.cards__container}>

        <div className={`${styles.cards__card} ${styles.color__primary}`}>
          <div className={styles.card__header}>
            <FaUser className={styles.header__icon}/>
            <p className={styles.header__text}>Usuarios</p>
          </div>
          <div className={styles.card__content}>
            <h2 className={styles.content__text}>
                100
            </h2>
          </div>
        </div>

        <div className={`${styles.cards__card} ${styles.color__primary}`}>
          <div className={styles.card__header}>
            <BsCashStack className={styles.header__icon}/>
            <p className={styles.header__text}>Recaudado</p>
          </div>
          <div className={styles.card__content}>
            <h2 className={styles.content__text}>
                <span>$</span>50,000
            </h2>
          </div>
        </div>

        <div className={`${styles.cards__card} ${styles.color__red}`}>
          <div className={styles.card__header}>
            <FaUserTimes className={styles.header__icon}/>
            <p className={styles.header__text}>Suspendidos</p>
          </div>
          <div className={styles.card__content}>
            <h2 className={styles.content__text}>
                100
            </h2>
          </div>
        </div>

        <div className={`${styles.cards__card} ${styles.color__red}`}>
          <div className={styles.card__header }>
            <BsCashStack className={styles.header__icon}/>
            <p className={styles.header__text}>Faltante</p>
          </div>
          <div className={styles.card__content}>
            <h2 className={styles.content__text}>
                <span>$</span>10,000
            </h2>
          </div>
        </div>

    </section>
  )
}

export default Cards