import Image from "next/image"
import { useRouter } from "next/router"
import { clsx } from "clsx"
import config from "@/utils/config"
import apiRoutes from "@/utils/constants"

const CategoryCard = ({ category, className, otherProps }) => {
  const router = useRouter()
  const goToCategoryPage = () => router.push(apiRoutes.categoryId(category.ID))

  return (
    <li className={clsx(className, "flex flex-col gap-1 max-w-md  rounded mx-auto")} onClick={goToCategoryPage} {...otherProps}>
      <Image
        alt={category.name}
        src={category.thumbnailURL}
        width={config.images.categoryCard.width}
        height={config.images.categoryCard.height}
        className="rounded-sm drop-shadow-box"
      />
      <div className="flex justify-between">
        <h2>{category.name}</h2>
      </div>
    </li>
  )
}

export default CategoryCard

