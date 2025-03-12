<script lang="ts">
    type OrderSubmitRequirementsInfoTitleProps = {
        title: string
    }

    const {title}: OrderSubmitRequirementsInfoTitleProps = $props();
</script>

<p class="font-semibold">{title}</p>