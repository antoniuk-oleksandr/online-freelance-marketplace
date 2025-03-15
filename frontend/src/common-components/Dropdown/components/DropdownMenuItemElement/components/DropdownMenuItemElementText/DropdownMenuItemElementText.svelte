<script lang="ts">
  type DropdownMenuItemElementTextProps = {
    title: string
  }

  const { title }: DropdownMenuItemElementTextProps = $props()
</script>

<span class="">{title}</span>
