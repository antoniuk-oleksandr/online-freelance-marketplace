<script lang="ts">
    type SuccessMessageTitleProps = {
        title: string,
    }

    const {title}: SuccessMessageTitleProps = $props();
</script>

<h1 class="text-2xl font-bold">{title}</h1>
