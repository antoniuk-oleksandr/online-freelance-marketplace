<script lang="ts">
  type OrderByIdSidebarServiceImageProps = {
    src: string
    alt: string
  }

  const { src, alt }: OrderByIdSidebarServiceImageProps = $props()
</script>

<img {src} {alt} class="aspect-video object-center object-cover rounded-sm w-32" />
