<script lang="ts">
    type OrderPageProps = {
        id: string,
    }

    const {id}: OrderPageProps = $props();
</script>

<div class="grid grid-cols-2">
    <div>
        <p>Some product info</p>
        <p>{id}</p>
        <textarea></textarea>
    </div>
    <div>

    </div>
</div>

