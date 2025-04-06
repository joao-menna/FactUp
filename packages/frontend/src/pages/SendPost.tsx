import { InsertPostInput, postService } from "services/post";
import TextAreaAutosize from "react-textarea-autosize";
import { TextField } from "lib/components/TextField";
import { useMutation } from "@tanstack/react-query";
import { FaRegCircleXmark } from "react-icons/fa6";
import { useTranslation } from "react-i18next";
import { Button } from "lib/components/Button";
import { imageService } from "services/image";
import { useDropzone } from "react-dropzone";
import { FaRegImage } from "react-icons/fa";
import { useNavigate } from "react-router";
import { useState } from "react";
import { clsx } from "clsx/lite";

type FormType = "fact" | "saying";

const handleDrop = (
  acceptedFiles: File[],
  callbackImageFile: (result: File) => void,
  callbackDataUrl: (result: string) => void
) => {
  if (acceptedFiles.length < 1) {
    return;
  }

  const file = acceptedFiles[0];

  const reader = new FileReader();
  reader.onload = (ev) => callbackDataUrl(ev.target?.result as string);

  reader.readAsDataURL(file);

  callbackImageFile(file);
};

export function SendPostPage() {
  const [currentDataUrl, setCurrentDataUrl] = useState<string | null>(null);
  const [selectedForm, setSelectedForm] = useState<FormType>("fact");
  const [imageFile, setImageFile] = useState<File | null>(null);
  const [source, setSource] = useState<string>();
  const [error, setError] = useState<string>("");
  const [body, setBody] = useState<string>("");
  const navigate = useNavigate();
  const { t } = useTranslation();

  const { getRootProps, getInputProps, isDragActive } = useDropzone({
    onDrop: (af) => handleDrop(af, setImageFile, setCurrentDataUrl),
    accept: {
      "image/gif": [".gif"],
      "image/jpeg": [".jpg", ".jpeg"],
      "image/png": [".png"],
      "image/webp": [".webp"],
    },
  });

  const imageMutation = useMutation({
    mutationFn: async (image: Blob) => await imageService.send(image),
  });

  const postMutation = useMutation({
    mutationFn: async (body: InsertPostInput) => await postService.insert(body),
  });

  const isMutationLoading = () =>
    imageMutation.isPending || postMutation.isPending;

  const handleChangeForm = (targetForm: "fact" | "saying") => {
    setSelectedForm(targetForm);
    setSource(undefined);
  };

  const handleSendPost = async () => {
    let treatedSource = source;
    let imagePath: string | undefined = undefined;

    if (!body && !imageFile) {
      setError(t("youNeedToInputABodyOrImage"));
      return;
    }

    if (imageFile) {
      try {
        imagePath = (await imageMutation.mutateAsync(imageFile)).imagePath;
      } catch {
        setError(t("youReachedTheDailyLimit"));
        return;
      }
    }

    if (!treatedSource) {
      treatedSource = undefined;
    }

    try {
      const post = await postMutation.mutateAsync({
        type: selectedForm,
        body: body ?? "",
        source,
        imagePath,
      });

      navigate(`/p/${post.id}`);
    } catch {
      setError(t("youReachedTheDailyLimit"));
    }
  };

  const handleClearImage = () => {
    setCurrentDataUrl(null);
    setImageFile(null);
  };

  return (
    <div className="flex flex-col justify-between h-full container">
      <div className={clsx("flex flex-col gap-2 p-2 pt-18")}>
        <div className="flex gap-2">
          <Button
            onClick={() => handleChangeForm("fact")}
            className={clsx(
              "w-full h-16",
              selectedForm === "fact"
                ? "bg-accent-400 hover:bg-accent-400/80"
                : "bg-accent-500 hover:bg-accent-500/80"
            )}
          >
            {t("fact")}
          </Button>
          <Button
            onClick={() => handleChangeForm("saying")}
            className={clsx(
              "w-full h-16",
              selectedForm === "saying"
                ? "bg-accent-400 hover:bg-accent-400/80"
                : "bg-accent-500 hover:bg-accent-500/80"
            )}
          >
            {t("saying")}
          </Button>
        </div>
        <div></div>
        <div>
          <div className={clsx("flex justify-between")}>
            <label className={clsx("text-text-200 text-lg")}>
              {selectedForm === "fact" ? t("fact") : t("saying")}
            </label>
            <span className="text-lg text-text-100">{280 - body.length}</span>
          </div>
          <TextAreaAutosize
            minRows={3}
            maxLength={280}
            value={body ?? ""}
            onChange={(e) => setBody(e.target.value)}
            className={clsx(
              "w-full border-2 border-accent-500 bg-primary-800",
              "rounded-lg p-1 outline-0 text-text-100"
            )}
          />
        </div>
        {selectedForm === "fact" && (
          <div>
            <div className="flex justify-between">
              <label className={clsx("text-text-200 text-lg")}>
                {t("source")}
              </label>
              <span className="text-lg text-text-100">
                {80 - (source?.length ?? 0)}
              </span>
            </div>
            <TextField
              type="text"
              maxLength={80}
              value={source ?? ""}
              onChange={(e) => setSource(e.target.value)}
              className={clsx(
                "w-full border-2 border-accent-500 bg-primary-800",
                "rounded-lg p-1 outline-0 text-text-100"
              )}
            />
          </div>
        )}
        {currentDataUrl ? (
          <div className="flex flex-col">
            <Button
              onClick={handleClearImage}
              className={clsx(
                "top-0 right-0 p-2 bg-accent-500",
                "hover:bg-accent-500/80 flex items-center",
                "gap-2 justify-between rounded-b-none"
              )}
            >
              <FaRegCircleXmark className="text-2xl" />
              <span className="text-lg">{t("removeImage")}</span>
              <FaRegCircleXmark className="text-2xl" />
            </Button>
            <img
              className={clsx("rounded-b-lg")}
              src={currentDataUrl}
              alt={t("imageToSend")}
            />
          </div>
        ) : (
          <div
            {...getRootProps()}
            className={clsx(
              "w-full border-2 flex flex-col justify-between rounded-lg",
              "border-accent-500 bg-primary-800 gap-2",
              "items-center"
            )}
          >
            <input {...getInputProps()} />
            <div
              className={clsx(
                "flex items-center justify-between text-text-200 size-full p-2",
                "bg-accent-500 hover:bg-accent-500/80 gap-4"
              )}
            >
              <FaRegImage className="text-lg" />
              <p className="text-text-200">
                {isDragActive
                  ? t("dropTheImageHere")
                  : t("dragYourImageHereOrClickToPickOne")}
              </p>
            </div>
          </div>
        )}
      </div>
      <div
        className={clsx(
          "absolute top-0 right-0 flex gap-2 justify-end",
          "items-center break-words whitespace-pre-wrap"
        )}
      >
        <p className="text-red-300 w-48">{error}</p>
        <Button
          onClick={handleSendPost}
          className={clsx("bg-accent-500 hover:bg-accent-500/80 m-2 h-14 w-24")}
          disabled={isMutationLoading()}
        >
          {isMutationLoading() ? t("loading") : t("post")}
        </Button>
      </div>
    </div>
  );
}
