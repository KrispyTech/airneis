import Image from "next/image"

const ProductCard = ({ product }) => (
  <li className="flex flex-col gap-1">
    <Image alt="#" src={product.image} width={450} height={200} className="rounded-sm drop-shadow-box" />
    <div className="flex justify-between">
      <h2>{product.name}</h2>
      <span>{product.price}€</span>
    </div>
  </li>
)

export default ProductCard
