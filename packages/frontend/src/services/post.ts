interface InsertPostInput {
  type: string;
  body: string;
  source: string;
  imagePath: string;
}

class PostService {
  baseUrl = `${import.meta.env.VITE_BACKEND_BASE_URL}/api/v1/post`;

  async findById(postId: number) {
    const req = await fetch(`${this.baseUrl}/single/${postId}`);

    const json = await req.json();

    return json;
  }

  async findRandom(type: string, limit: number) {
    const req = await fetch(
      `${this.baseUrl}/multiple/random?type=${type}&limit=${limit}`
    );

    const json = await req.json();

    return json;
  }

  async findPagedByUser(userId: number, limit: number, page: number) {
    const req = await fetch(
      `${this.baseUrl}/multiple/user/${userId}?limit=${limit}&page=${page}`
    );

    const json = await req.json();

    return json;
  }

  async insertPost(body: InsertPostInput) {
    const req = await fetch(this.baseUrl, {
      method: "POST",
      body: JSON.stringify(body),
    });

    const json = await req.json();

    return json;
  }

  async delete(postId: number) {
    const req = await fetch(`${this.baseUrl}/${postId}`, {
      method: "DELETE",
    });

    const json = await req.json();

    return json;
  }
}

export const imageService = new PostService();
