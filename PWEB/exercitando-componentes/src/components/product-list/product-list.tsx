import products from "../../mock";
import CardProduct from "@/components/card-product.tsx/card-product";
import styles from "./product-list.module.css";
export default function ProductList() {
  return (
    <div className={styles.productListRow}>
      {products.map((product) => (
        <CardProduct
          key={product.id}
          id={product.id}
          name={product.name}
          description={product.description}
          price={product.price}
          imageSRC={product.imageSRC}
        />
      ))}
    </div>
  );
}
