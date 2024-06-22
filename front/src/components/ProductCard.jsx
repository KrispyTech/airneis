import Image from "next/image"
import { useRouter } from "next/router"
import { clsx } from "clsx"
import config from "@/utils/config"

const ProductCard = ({ product, className, otherProps }) => {
  const router = useRouter()
  const goToProductPage = () => router.push(`/products/${product.ID}`)

  return (
    <li className={clsx(className, "flex flex-col gap-1 max-w-md  rounded mx-auto")} onClick={goToProductPage} {...otherProps}>
      <Image
        alt={product.name}
        src={product.thumbnailURL}
        width={config.images.productCard.imageWidth}
        height={config.images.productCard.imageHeight}
        className="rounded-sm drop-shadow-box"
      />
      <div className="flex justify-between">
        <h2>{product.name}</h2>
        <span>{product.priceWithTaxes}€</span>
      </div>
    </li>
  )
}

export default ProductCard
