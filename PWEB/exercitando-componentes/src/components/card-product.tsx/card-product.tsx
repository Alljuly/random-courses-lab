import styles from "./card-product.module.css";

interface Props {
  id: number;
  name: string;
  description: string;
  price: string;
  imageSRC: string;
}

export default function CardProduct({
  id,
  name,
  description,
  price,
  imageSRC,
}: Props) {
  return (
    <div className={styles.cardProduct} id={`${id}`}>
      <img className={styles.cardImage} src={imageSRC} alt="..." />
      <div className={styles.atributtesProduct}>
        <p className="name-product">{name}</p>
        <p>{description}</p>
        <p className="price-product">{price}</p>
      </div>
    </div>
  );
}
