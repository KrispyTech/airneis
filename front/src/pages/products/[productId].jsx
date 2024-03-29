import Page from "@/components/Page"
import Button from "@/components/ui/Button"
import MaterialChip from "@/components/materials/MaterialChip"
import { getProductById } from "@/utils/services/productServices"

export const getServerSideProps = async ({ query }) => {
  const { productId } = query
  const { data: product } = await getProductById(productId)

  return { props: { product } }
}
const ProductPage = props => {
  const {
    product: { name, description, outOfStock, priceWithTaxes, materials }
  } = props

  return (
    <Page>
      <section className={"bg-blue-500 h-[250px] flex items-center justify-center"}>
        <h1 className="text-4xl text-white font-semibold">{name}</h1>
      </section>
      <div className="mx-12 mt-8 ">
        <section className="flex flex-col gap-7 w-full mt-3 sm:flex-row">
          <div className="sm:w-[400px] flex justify-center items-center border-2 border-black">Carrousel</div>
          <div>
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
              <Button text="Ajouter au panier" variant="orange"></Button>
              <div className="flex flex-row gap-5 mt-5">
                {materials && materials.map((material, index) => <MaterialChip name={material.name} key={index} />)}
              </div>
            </div>
          </div>
        </section>
        <section className="mt-7">
          <h2 className="text-center font-bold text-xl">PRODUITS SIMILAIRES</h2>
          <div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 gap-5 mt-5">
            <div className="border-2 border-black">Produit 1</div>
            <div className="border-2 border-black">Produit 2</div>
            <div className="border-2 border-black">Produit 3</div>
            <div className="border-2 border-black">Produit 4</div>
            <div className="border-2 border-black">Produit 5</div>
            <div className="border-2 border-black">Produit 6</div>
          </div>
        </section>
      </div>
    </Page>
  )
}

export default ProductPage
