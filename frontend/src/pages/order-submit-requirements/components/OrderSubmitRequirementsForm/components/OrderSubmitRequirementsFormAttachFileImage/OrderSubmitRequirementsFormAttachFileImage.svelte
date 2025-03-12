<script lang="ts">
  type OrderSubmitRequirementsFormAttachFileImageProps = {
    file: File
  }

  const { file }: OrderSubmitRequirementsFormAttachFileImageProps = $props()
</script>

<img
  class="object-center object-cover w-full h-full rounded-xl"
  src={URL.createObjectURL(file)}
  alt={file.name}
/>
