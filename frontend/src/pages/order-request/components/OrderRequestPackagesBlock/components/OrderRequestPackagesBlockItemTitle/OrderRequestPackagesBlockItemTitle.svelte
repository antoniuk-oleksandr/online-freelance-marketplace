<script lang="ts">
    type OrderRequestPackagesBlockItemTitleProps = {
        title: string,
    };

    const {title}: OrderRequestPackagesBlockItemTitleProps = $props();
</script>

<div class="font-medium text-lg ">{title}</div>
