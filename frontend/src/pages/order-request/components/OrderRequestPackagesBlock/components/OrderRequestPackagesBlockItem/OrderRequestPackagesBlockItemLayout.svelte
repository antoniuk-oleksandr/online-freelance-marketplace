<script lang="ts">
  import type { LayoutProps } from '@/types/LayoutProps.ts'
  import type { Package } from '@/types/Package.ts'

  type OrderRequestPackagesBlockProps = LayoutProps & {
    setSelectedPackage: (newPackage: Package) => void
    package: Package
  }

  const { setSelectedPackage, children, package: pkg }: OrderRequestPackagesBlockProps = $props()
</script>

<button
  type="button"
  onclick={() => setSelectedPackage(pkg)}
  class="flex text-start items-center cursor-pointer gap-3 p-3 rounded-xl duration-200 ease-out hover:bg-light-palette-action-hover hover:dark:bg-dark-palette-action-hover"
>
  {@render children()}
</button>
