import "./app-module.css";
import "./globals.css";
import Header from "@/components/header/header";
import ProductList from "@/components/product-list/product-list";

export default function Home() {
  return (
    <div className="Home">
      <Header />
      <main className="content">
        <img src="/assets/jewelry-header.webp" className="banner" alt="" />
        <div className="container-products">
          <ProductList />
        </div>
      </main>
      <footer>oiiii</footer>
    </div>
  );
}
