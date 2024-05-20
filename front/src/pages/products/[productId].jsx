import Page from "@/components/Page"
import Button from "@/components/ui/buttons/Button"
import MaterialChip from "@/components/materials/MaterialChip"
import { getProductById } from "@/utils/services/productServices"
import Carousel from "@/components/Carousel"
import ProductCard from "@/components/ProductCard"
import { serverSideTranslations } from "next-i18next/serverSideTranslations"
import { useTranslation } from "next-i18next"

export const getServerSideProps = async ({ query, locale }) => {
  const { productId } = query
  const { data: product } = await getProductById(productId)
  // AS MATERIALS ARE NOT IMPLEMENTED BACKEND SIDE, WE ADD THEM MANUALLY
  product.materials = [{ name: "wood" }, { name: "leather" }, { name: "cotton" }]

  return { props: { product, ...(await serverSideTranslations(locale, ["common"])) } }
}
const ProductPage = ({ product: { name, description, outOfStock, priceWithTaxes, materials }, product }) => {
  const { t } = useTranslation("common")

  return (
    <Page>
      <section
        className={"h-[250px] hidden md:flex items-center justify-center bg-cover bg-center bg-no-repeat bg-[image:var(--image-url)] "}
        style={{ "--image-url": `url(${product.thumbnailURL})` }}
      >
        <h1 className="text-4xl text-white font-semibold ">{name}</h1>
      </section>
      <div className="md:mx-12 md:mt-8">
        <section className="flex flex-col gap-7 w-full md:mt-3 lg:flex-row ">
          <Carousel
            images={["https://picsum.photos/452/250", "https://picsum.photos/453/250", "https://picsum.photos/454/250"]}
            text={product.name}
            className="w-full"
          />
          <div className="w-full px-8 ">
            <div>
              <div className="flex flex-row justify-between font-bold text-2xl">
                <h3>{priceWithTaxes} €</h3>
                <h2>{name}</h2>
              </div>
              <div className="flex md:justify-end text-slate-500 mt-3 text-lg">
                <p>{outOfStock ? t("productPage.outOfStock") : t("productPage.inStock")}</p>
              </div>
            </div>
            <div className="md:mt-5 font-semibold mt-5">
              <p>{description}</p>
            </div>
            <div className="flex flex-col itens-center xl:flex-row justify-between mt-5">
              <Button disabled={outOfStock} text={t("productPage.addToCart")} variant="orange" />
              <div className="flex flex-row items-center gap-5 mt-5 h-full">
                {materials && materials.map((material, index) => <MaterialChip name={material.name} key={index} />)}
              </div>
            </div>
          </div>
        </section>
        <section className="mt-7">
          <h2 className="text-center font-extrabold text-4xl uppercase">{t("productPage.similarProducts")}</h2>
          <div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 gap-x-5 gap-y-12 cols-gap-2 mt-12 mx-4">
            {
              // Need to be replaced when backend will send similar products
            }
            <ProductCard product={product} />
          </div>
        </section>
      </div>
    </Page>
  )
}

export default ProductPage
