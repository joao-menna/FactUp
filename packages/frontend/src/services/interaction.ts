class InteractionService {
  baseUrl = `${import.meta.env.VITE_BACKEND_BASE_URL}/api/v1/interaction`;

  async get(postId: number) {
    const req = await fetch(`${this.baseUrl}/${postId}`);

    const json = await req.json();

    return json;
  }

  async addPositive(postId: number) {
    const req = await fetch(this.baseUrl, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ postId, score: 1 }),
    });

    if (req.status !== 200) {
      throw new Error("Could not add positive vote");
    }

    const json = await req.json();

    return json;
  }

  async addNegative(postId: number) {
    const req = await fetch(this.baseUrl, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ postId, score: -1 }),
    });

    if (req.status !== 200) {
      throw new Error("Could not add negative vote");
    }

    const json = await req.json();

    return json;
  }

  async remove(postId: number) {
    const req = await fetch(`${this.baseUrl}?postId=${postId}`, {
      method: "DELETE",
    });

    if (req.status !== 200) {
      throw new Error("Could not add negative vote");
    }

    const json = await req.json();

    return json;
  }
}

export const interactionService = new InteractionService();
