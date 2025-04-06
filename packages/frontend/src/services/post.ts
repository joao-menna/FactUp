export interface InsertPostInput {
  type: string;
  body: string;
  source?: string;
  imagePath?: string;
}

export interface Post {
  id: number;
  type: "fact" | "saying";
  userId: number;
  body: string;
  source?: string;
  imagePath?: string;
  createdAt?: string;
}

class PostService {
  baseUrl = `${import.meta.env.VITE_BACKEND_BASE_URL ?? ""}/api/v1/post`;

  async findById(postId: number) {
    const req = await fetch(`${this.baseUrl}/single/${postId}`, {
      credentials: "include",
    });

    if (req.status !== 200) {
      throw new Error("could not find post");
    }

    const json = (await req.json()) as Post;

    return json;
  }

  async findPaged(type: string, limit: number, page: number) {
    const req = await fetch(
      `${this.baseUrl}/multiple?type=${type}&limit=${limit}&page=${page}`,
      {
        credentials: "include",
      }
    );

    const json = (await req.json()) as Post[];

    return json;
  }

  async findRandom(type: string, limit: number) {
    const req = await fetch(
      `${this.baseUrl}/multiple/random?type=${type}&limit=${limit}`,
      {
        credentials: "include",
      }
    );

    const json = (await req.json()) as Post[];

    return json;
  }

  async findPagedByUser(userId: number, limit: number, page: number) {
    const req = await fetch(
      `${this.baseUrl}/multiple/user/${userId}?limit=${limit}&page=${page}`,
      {
        credentials: "include",
      }
    );

    const json = (await req.json()) as Post[];

    return json;
  }

  async insert(body: InsertPostInput) {
    const req = await fetch(this.baseUrl, {
      method: "POST",
      body: JSON.stringify(body),
      credentials: "include",
    });

    if (req.status !== 200) {
      throw new Error("reached maximum post count for this account today");
    }

    const json = (await req.json()) as Post;

    return json;
  }

  async delete(postId: number) {
    const req = await fetch(`${this.baseUrl}/${postId}`, {
      method: "DELETE",
      credentials: "include",
    });

    const json = await req.json();

    return json;
  }
}

export const postService = new PostService();
