import { Image, TouchableHighlight, View } from "react-native"
import mobileConfig from "../../mobileConfig"

const ImageGallery = props => {
  const { images, imageToDisplay, handleChangeImage } = props

  return (
    <View>
      <Image style={mobileConfig.images.fullScreenImageLG} className="rounded-b-md drop-shadow-box" source={{ url: images[imageToDisplay] }} />
      <View className="flex-row justify-center gap-4 pt-2">
        {images.map((_, i) => (
          <TouchableHighlight key={`image-${i}`} onPress={handleChangeImage(i)} data-pos={i}>
            <View className={`border-y-6 border-x-6 w-6 h-6 rounded-full ${imageToDisplay === i ? "bg-green-400" : ""}`} />
          </TouchableHighlight>
        ))}
      </View>
    </View>
  )
}

export default ImageGallery
