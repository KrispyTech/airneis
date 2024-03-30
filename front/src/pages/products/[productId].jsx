import Page from "@/components/Page"
import Button from "@/components/ui/Button"
import MaterialChip from "@/components/materials/MaterialChip"
import { getProductById } from "@/utils/services/productServices"
import Carousel from "@/components/Carousel"
import ProductCard from "@/components/ProductCard"

export const getServerSideProps = async ({ query }) => {
  const { productId } = query
  const { data: product } = await getProductById(productId)

  return { props: { product } }
}
const ProductPage = ({ product: { name, description, outOfStock, priceWithTaxes, materials }, product }) => (
  <Page>
    <section
      className={"h-[250px] hidden sm:flex items-center justify-center bg-cover bg-center bg-no-repeat bg-[image:var(--image-url)] "}
      style={{ "--image-url": `url(${product.thumbnailURL})` }}
    >
      <h1 className="text-4xl text-white font-semibold ">{name}</h1>
    </section>
    <div className="sm:mx-12 sm:mt-8">
      <section className="flex flex-col gap-7 w-full sm:mt-3 sm:flex-row">
        <Carousel images={["https://picsum.photos/452/250", "https://picsum.photos/453/250", "https://picsum.photos/454/250"]} text={product.name} />
        <div className="w-full px-8 ">
          <div>
            <div className="flex flex-row  justify-between font-bold text-xl">
              <h3>{priceWithTaxes} €</h3>
              <h2>{name}</h2>
            </div>
            <div className="flex sm:justify-end">
              <p>{outOfStock ? "Stocks épuisés" : "En stock"}</p>
            </div>
          </div>
          <div className="sm:mt-5 mt-2">
            <p>{description}</p>
          </div>
          <div className="flex flex-col sm:flex-row justify-between mt-3">
            <Button disabled={outOfStock} text="Ajouter au panier" variant="orange"></Button>
            <div className="flex flex-row gap-5 mt-5">
              {materials && materials.map((material, index) => <MaterialChip name={material.name} key={index} />)}
            </div>
          </div>
        </div>
      </section>
      <section className="mt-7">
        <h2 className="text-center font-bold text-xl">PRODUITS SIMILAIRES</h2>
        <div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 gap-x-5 gap-y-12 cols-gap-2 mt-12">
          <ProductCard product={product} />
          <ProductCard product={product} />
          <ProductCard product={product} />
          <ProductCard product={product} />
          <ProductCard product={product} />
          <ProductCard product={product} />
        </div>
      </section>
    </div>
  </Page>
)

export default ProductPage
