import MyForm from "@/components/ui/forms/MyForm"
import SubmitButton from "@/components/ui/buttons/SubmitButton"
import FormField from "@/components/ui/forms/FormField"
import FormFieldSelect from "@/components/ui/forms/FormFieldSelect"
import UploadAndDisplayImage from "@/components/ui/forms/UploadAndDisplayImages"

const AddProductForm = ({ images, setImages, categories, t }) => (
  <MyForm className="text-2xl">
    <h1 className="font-semibold text-3xl">NEW</h1>
    <div className="flex justify-between gap-x-16">
      <div className="flex flex-col items-center justify-center">
        <UploadAndDisplayImage images={images} setImages={setImages} />
      </div>
      <div className="flex flex-1 flex-col justify-start ">
        <FormField name="name" type="text" placeholder={t("backoffice.placeholder.name")} label={t("backoffice.label.name") + "*"} className="py-2" />
        <FormField
          name="orderOfPriority"
          type="number"
          placeholder={t("backoffice.placeholder.orderOfPriority")}
          label={t("backoffice.label.orderOfPriority")}
          className="py-2"
        />
        <FormField
          name="priceWithoutTaxes"
          type="number"
          placeholder={t("backoffice.placeholder.priceWithoutTaxes")}
          label={t("backoffice.label.priceWithoutTaxes") + "*"}
          className="py-2"
        />
        <FormField
          name="priceWithTaxes"
          type="number"
          placeholder={t("backoffice.placeholder.priceWithTaxes")}
          label={t("backoffice.label.priceWithTaxes") + "*"}
          className="py-2"
        />
        <FormField
          name="materialsString"
          type="text"
          placeholder={t("backoffice.placeholder.materials")}
          label={t("backoffice.label.materials") + "*"}
          className="py-2"
        />
        <FormFieldSelect
          name="categoryID"
          placeholder={t("backoffice.placeholder.category")}
          label={t("backoffice.label.category")}
          options={categories}
        />
      </div>
    </div>
    <FormField name="description" as="textarea" placeholder={t("backoffice.placeholder.description")} label={t("backoffice.label.description")} />
    <div className="">
      <SubmitButton text={"validate"} className="mt-4 py-0.5 px-2.5 bg-gray w-fit rounded-full border flex gap-1 items-center" />
    </div>
  </MyForm>
)

export default AddProductForm
