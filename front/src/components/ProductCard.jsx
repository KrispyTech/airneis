import Image from "next/image"
import { useRouter } from "next/router"

const ProductCard = ({ product }) => {
  const router = useRouter()
  const goToProductPage = () => router.push(`/products/${product.ID}`)

  return (
    <li className="flex flex-col gap-1 max-w-md border border-secondary p-2 rounded" onClick={goToProductPage}>
      <Image alt={product.name} src={product.image} width={450} height={200} className="rounded-sm drop-shadow-box" />
      <div className="flex justify-between">
        <h2>{product.name}</h2>
        <span>{product.price}€</span>
      </div>
    </li>
  )
}

export default ProductCard
