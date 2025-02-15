export namespace internal {
	
	export class TextureFile {
	    name: string;
	    imgData: string;
	    width: number;
	    height: number;
	
	    static createFrom(source: any = {}) {
	        return new TextureFile(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.imgData = source["imgData"];
	        this.width = source["width"];
	        this.height = source["height"];
	    }
	}

}

