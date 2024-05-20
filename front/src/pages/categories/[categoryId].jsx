/* eslint-disable no-unused-vars */
import Page from "@/components/Page"
import ProductCard from "@/components/ProductCard"
import { getCategoryById } from "@/utils/services/categoryServices"
import { serverSideTranslations } from "next-i18next/serverSideTranslations"

export const getServerSideProps = async ({ query, locale }) => {
  const { categoryId } = query
  const { data: category } = await getCategoryById(categoryId)

  return { props: { category, ...(await serverSideTranslations(locale, ["common"])) } }
}
const CategoryPage = ({ category }) => (
  <Page>
    <div className="w-full h-1/4 sm:h-1/3 bg-[url('https://picsum.photos/1800/350')] bg-center flex items-center justify-center">
      <h1 className="text-white drop-shadow-textOutline font-bold text-xl lg:text-3xl">{category.name}</h1>
    </div>
    <p className="font-semibold text-3xl px-8 py-14">{category.description}</p>
    <ul className="grid grid-cols-3 gap-x-12 gap-y-8 px-8 pb-6">
      {category.products.map(product => (
        <ProductCard key={product.ID} product={product} />
      ))}
    </ul>
  </Page>
)

export default CategoryPage
