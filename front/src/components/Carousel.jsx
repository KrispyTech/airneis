import { useState } from "react"
import Image from "next/image"
import { clsx } from "clsx"
import config from "@/utils/config"

const CarouselControllerIcon = ({ filled }) => {
  if (filled) {
    return (
      <svg width="22" height="21" viewBox="0 0 22 21" fill="none" xmlns="http://www.w3.org/2000/svg">
        <path
          d="M19.5 10.5C19.5 14.81 15.8053 18.5 11 18.5C6.19471 18.5 2.5 14.81 2.5 10.5C2.5 6.19005 6.19471 2.5 11 2.5C15.8053 2.5 19.5 6.19005 19.5 10.5Z"
          fill="#A1D399"
          stroke="#002203"
          strokeWidth="5"
        />
      </svg>
    )
  }

  return (
    <svg width="22" height="21" viewBox="0 0 22 21" fill="none" xmlns="http://www.w3.org/2000/svg">
      <path
        d="M19.5 10.5C19.5 14.81 15.8053 18.5 11 18.5C6.19471 18.5 2.5 14.81 2.5 10.5C2.5 6.19005 6.19471 2.5 11 2.5C15.8053 2.5 19.5 6.19005 19.5 10.5Z"
        stroke="#002203"
        strokeWidth="5"
      />
    </svg>
  )
}

const Carousel = ({ images, className, text, ...otherProps }) => {
  const [current, setCurrent] = useState(0)
  const handleChangeImage = index => () => {
    setCurrent(index)
  }

  return (
    <div className={clsx("flex flex-col items-center justify-center", className)} {...otherProps}>
      <div className="relative w-full flex justify-center">
        {images &&
          images.map((image, index) => (
            <Image
              alt={`image-${index}`}
              key={index}
              src={image}
              width={config.images.productPage.carousel.width}
              height={config.images.productPage.carousel.height}
              className={`rounded-sm drop-shadow-box w-full lg:w-none ${index !== current ? "hidden" : ""}`}
            />
          ))}
        <h2 className="md:hidden absolute top-1/2 left-1/2 -translate-y-1/2 -translate-x-1/2 text-4xl text-white font-semibold whitespace-nowrap drop-shadow-box">
          {text}
        </h2>
      </div>
      <div className="flex justify-evenly mt-5 w-48">
        {images.map((_, index) => (
          <button key={index} onClick={handleChangeImage(index)}>
            <CarouselControllerIcon filled={index === current} />
          </button>
        ))}
      </div>
    </div>
  )
}

export default Carousel
