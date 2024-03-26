/* eslint-disable no-unused-vars */
import Page from "@/components/Page"
import ProductCard from "@/components/ProductCard"

const category = {
  description: "LES SALONS ÉCOSSAIS OFFRENT UNE POSSIBILITÉ INFINIE À VOTRE MAISON. DÉCOUVREZ LA DOUCEUR ET LA FRAÎCHEUR DE VOTRE FUTUR SALON."
}

const products = [
  { id: 1, name: "CANAPÉ BLANC", price: 1200, image: "https://picsum.photos/450/250" },
  { id: 2, name: "CANAPÉ SAUFA", price: 2450, image: "https://picsum.photos/450/250" },
  { id: 3, name: "CHAISE DES LANDES", price: 575, image: "https://picsum.photos/450/250" },
  { id: 4, name: "CANAPÉ BLANC", price: 1200, image: "https://picsum.photos/450/250" },
  { id: 5, name: "CANAPÉ SAUFA", price: 2450, image: "https://picsum.photos/450/250" },
  { id: 6, name: "CHAISE DES LANDES", price: 575, image: "https://picsum.photos/450/250" }
]

export const getServerSideProps = ({ query }) => {
  const { categoryId } = query

  return { props: { categoryId } }
}
const CategoryPage = props => {
  const { categoryId } = props

  return (
    <Page>
<<<<<<< HEAD
      <div className="w-full h-1/4 sm:h-1/3 bg-[url('https://picsum.photos/1800/350')] bg-center flex items-center justify-center">
        <h1 className="text-white drop-shadow-textOutline font-bold text-xl lg:text-3xl">SALON</h1>
      </div>
=======
      <button className="w-full h-1/4 sm:h-1/3 bg-[url('https://picsum.photos/1800/350')] bg-center flex items-center justify-center">
        <h1 className="text-white drop-shadow-[0_0_1px_rgba(0,0,0,1)] font-bold text-xl lg:text-3xl">SALON</h1>
      </button>
>>>>>>> 314558b (fix: separating PR)
      <p className="font-semibold text-3xl px-8 py-14">{category.description}</p>
      <ul className="grid grid-cols-3 gap-x-12 gap-y-8 px-8 pb-6">
        {products.map(product => (
          <ProductCard key={product.id} product={product} />
        ))}
      </ul>
    </Page>
  )
}

export default CategoryPage
