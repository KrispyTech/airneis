import Input from "../Input"
import Carousel from "@/components/Carousel"
import config from "@/utils/config"
import { PhotoIcon } from "@heroicons/react/24/outline"

const UploadAndDisplayImages = ({ images, setImages }) => {
  const handleAddImages = event => {
        setImages(() => [...event.target.files].map(image => URL.createObjectURL(image)))
      }
  
  return (
    <div className="flex flex-col items-center">
      <Carousel images={images} width={config.images.backoffice.carousel.width} height={config.images.backoffice.carousel.height} />
      <label htmlFor="file-upload" type="button" className="mt-4 py-0.5 px-3.5 bg-gray w-fit rounded-full border flex gap-1 items-center">
        <PhotoIcon className="w-6 h-6" />
        Add
      </label>
      <Input
        id="file-upload"
        type="file"
        name="myImage"
        className={"hidden"}
        multiple
        onChange={handleAddImages}
      />
    </div>
  )}

export default UploadAndDisplayImages
