import React, { useState, Im } from "react";
import Input from "../Input";
import Image from "next/image"
import Carousel from "@/components/Carousel";
import config from "@/utils/config";

const UploadImages = () => {
  const [images, setImages] = useState(["/no_img.jpeg"]);

  console.log(images)

  return (
    <div className="flex flex-col items-center">
      <Carousel 
        images={images}
        width={config.images.backoffice.carousel.width}
        height={config.images.backoffice.carousel.height}
      />
      <label for="file-upload" className="mt-4 py-0.5 px-2.5 bg-gray w-fit rounded-full border flex gap-1 items-center">
        Add
      </label>
        <Input
          id="file-upload"
          type="file"
          name="myImage"
          className={"hidden"}
          multiple
          onChange={(event) => {
            setImages(() => [...event.target.files].map((image) => URL.createObjectURL(image)));
            }}
        />
    </div>
  );
};

export default UploadImages;