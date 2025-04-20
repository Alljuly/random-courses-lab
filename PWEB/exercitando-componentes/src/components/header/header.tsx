import styles from "./header.module.css";

export default function Header() {
  return (
    <header className={styles.headerComponent}>
      <div className={styles.searchBar}>
        <p>busca</p>
        <p>logo</p>
        <p>perfil</p>
      </div>
      <div className={styles.navBar}>
        <p>home</p>
        <p>Produtos</p>
        <p>Contato</p>
        <p>Quem Somos</p>
        <p>Informações e Cuidados</p>
      </div>
    </header>
  );
}
