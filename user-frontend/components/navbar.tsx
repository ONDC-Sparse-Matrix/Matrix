"use-client";

import { Button } from "./ui/button";
import { Plus, Upload } from "lucide-react";
import Link from "next/link";
import {
  Tooltip,
  TooltipContent,
  TooltipProvider,
  TooltipTrigger,
} from "@/components/ui/tooltip";

import { UploadCSV } from "./upload-csv";

export function Navbar() {
  return (
    <>
      <nav className="bg-transparent w-full fixed backdrop-blur-lg dark:bg-gray-900">
        <div className="max-w-screen-xl flex flex-wrap items-center justify-between mx-auto p-4">
          <a
            href="/"
            className="flex items-center space-x-3 rtl:space-x-reverse"
          >
            <span className="text-2xl font-extrabold flex text-blue-500 whitespace-nowrap dark:text-white">
              matriX
            </span>
          </a>
          <div className="hidden w-full md:block md:w-auto" id="navbar-default">
            <TooltipProvider>
              <ul className="p-4 md:p-0 mt-4 flex rounded-lg bg-gray-50 md:flex-row md:space-x-1 rtl:space-x-reverse md:mt-0 md:border-0 md:bg-white dark:bg-gray-800 md:dark:bg-gray-900 dark:border-gray-700">
                <li>
                  <Link href="/addMerchant">
                    <Tooltip>
                      <TooltipTrigger asChild>
                        <Button variant={"link"}>
                          <Plus />
                        </Button>
                      </TooltipTrigger>
                      <TooltipContent>
                        <p>Add Merchant</p>
                      </TooltipContent>
                    </Tooltip>
                  </Link>
                </li>
                <li>
                  <UploadCSV />
                </li>
              </ul>
            </TooltipProvider>
          </div>
        </div>
      </nav>
    </>
  );
}
