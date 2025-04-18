interface PostImageOutput {
  imagePath: string;
  message: string;
}

class ImageService {
  baseUrl = `${import.meta.env.VITE_BACKEND_BASE_URL ?? ""}/api/v1/image`;

  async send(blob: Blob) {
    const body = new FormData();

    body.append("image", blob);

    const req = await fetch(this.baseUrl, {
      credentials: "include",
      method: "POST",
      body,
    });

    if (req.status === 413) {
      throw new Error("image too large");
    }

    if (req.status !== 200) {
      throw new Error("could not send image");
    }

    const json = (await req.json()) as PostImageOutput;

    return json;
  }
}

export const imageService = new ImageService();
