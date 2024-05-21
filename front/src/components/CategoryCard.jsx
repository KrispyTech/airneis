import Image from "next/image"
import { useRouter } from "next/router"
import { clsx } from "clsx"
import config from "@/utils/config"
import routes from "@@/lib/constants/routes"

const CategoryCard = ({ category, className, otherProps }) => {
  const router = useRouter()
  const goToCategoryPage = () => router.push(routes.categoryId(category.ID))

  return (
    <li className={clsx(className, "flex flex-col gap-1 max-w-md  rounded mx-auto relative")} onClick={goToCategoryPage} {...otherProps}>
      <Image
        alt={category.name}
        src={category.thumbnailURL}
        width={config.images.categoryCard.imageWidth}
        height={config.images.categoryCard.imageHeight}
        className="rounded-sm drop-shadow-box"
      />
      <h2 className="absolute top-1/2 left-1/2 -translate-x-1/2 -translate-y-1/2 text-white drop-shadow-textOutline font-bold text-xl">
        {category.name}
      </h2>
    </li>
  )
}

export default CategoryCard
