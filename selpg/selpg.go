package main

import(
	"fmt"
	"os"
	"os/exec"
	"bufio"
	"io"
	flag "github.com/spf13/pflag"
)

type selpg_args struct{
	start_page int
	end_page int
	page_len int 
	page_type string  
	in_filename string
	print_dest string
}

type sp_args selpg_args

var progname string

func main(){
	var sa sp_args
	progname = os.Args[0]
	
	initPflag(&sa)
	process_args(&sa)
	process_input(sa)
}

func usage(){
	fmt.Printf("\nUsage: \n\n")
	fmt.Printf("\t%s -sNumber -eNumber [ -f | -lNumber ] [ -dDestination] [ filename ]\n\n", progname)
	fmt.Printf("\t-sNumber\tStart from Page <Number>.\n")
	fmt.Printf("\t-eNumber\tEnd to Page <Number>.\n")
	fmt.Printf("\t-lNumber\tSpecify the number of line per page.Default is 72.\n")
	fmt.Printf("\t-f\t\t'f' for form-feed-delimited.Default is 'l'.(lines-delimited,)\n")
	fmt.Printf("\t-dDestination\t printer\n")
	fmt.Printf("\tfilename\tRead input from the file.Default stdin.(Control-D to end.)\n\n")
}

func initPflag(psa *sp_args){
	flag.Usage = usage
	flag.IntVarP(&psa.start_page, "start", "s", -1, "start page")
	flag.IntVarP(&(psa.end_page), "end", "e", -1, "end page")
	flag.IntVarP(&(psa.page_len), "len", "l", 72, "line number per")
	flag.StringVarP(&(psa.page_type), "type", "f","l", "'l' for lines-delimited, 'f' for form-feed-delimited. default is 'l'" )
	flag.StringVarP(&(psa.print_dest), "dest", "d", "", "print dest")
	flag.Lookup("type").NoOptDefVal = "f"
	flag.Parse()
}

func process_args(psa *sp_args) {
	if len(os.Args) < 3 {
		fmt.Fprintf(os.Stderr, "%s: not enough arguments\n", progname)
		flag.Usage()
		os.Exit(1)
	}

	if os.Args[1][0] != '-' || os.Args[1][1] != 's' {
		fmt.Fprintf(os.Stderr, "%s: 1st arg should be -sNumber\n", progname)
		flag.Usage()
		os.Exit(1)
	}

	INT_MAX := int(^uint(0) >> 1)

	if psa.start_page < 1 || psa.start_page > (INT_MAX - 1) {
		fmt.Fprintf(os.Stderr, "%s: invalid start page %d\n", progname, psa.start_page)
		flag.Usage()
		os.Exit(1)
	}

	if os.Args[2][0] != '-' || os.Args[2][1] != 'e' {
		fmt.Fprintf(os.Stderr, "%s: 2nd arg should be -eNumber\n", progname)
		flag.Usage()
		os.Exit(1)
	}

	if psa.end_page < 1 || psa.end_page > (INT_MAX - 1) ||  psa.end_page < psa.start_page{
		fmt.Fprintf(os.Stderr, "%s: invalid end page %d\n", progname, psa.end_page)
		flag.Usage()
		os.Exit(1)
	}


	if psa.page_len < 1 || psa.page_len > INT_MAX - 1 {
		fmt.Fprintf(os.Stderr, "%s: invalid page length %d\n", progname, psa.page_len)
		flag.Usage()
		os.Exit(1)
	}


	argLeft := flag.Args()
	if len(argLeft) > 0{
		_, err  := os.Stat(flag.Args()[0])
		if err != nil && os.IsNotExist(err){
			fmt.Fprintf(os.Stderr, "%s: the file <%s> does not exist, please create it first\n", progname, flag.Args()[0])
			os.Exit(1)
		}
		psa.in_filename  = string(argLeft[0])
	}else{
		psa.in_filename = ""
	}
}

func process_input(sa sp_args){
	cmd := &exec.Cmd{}

	//input
	var fin *os.File
	if len(sa.in_filename) == 0{
		fin = os.Stdin
	}else{
		var err error
		fin, err = os.Open(sa.in_filename)
		if err != nil{
			fmt.Fprintf(os.Stderr, "\n%s: could not open input file <%s>", progname, sa.in_filename)
			os.Exit(1)
		}
	}
	bufFin := bufio.NewReader(fin)

	//output
	var fout io.WriteCloser
	if len(sa.print_dest) == 0{
		fout = os.Stdout
	}else{
		cmd = exec.Command("cat") 
		var err error
		cmd.Stdout, err = os.OpenFile(sa.print_dest, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		if err != nil {
			fmt.Fprintf(os.Stderr, "\n%s: could not open file %s\n",
				progname, sa.print_dest)
			os.Exit(1)
		}
		
		//StdinPipe返回一个连接到command标准输入的管道pipe 
		fout, err = cmd.StdinPipe()
		if err != nil {
			fmt.Fprintf(os.Stderr, "\n%s: could not open pipe to \"lp -d%s\"\n",
				progname, sa.print_dest)
			os.Exit(1)
		}
		cmd.Start()
		defer fout.Close()
	}

	var line_ctr int
	var page_ctr int
	
	if sa.page_type == "l"{
		line_ctr = 0
		page_ctr = 1
		
		for {
			line, crc := bufFin.ReadString('\n')
			if crc != nil{
				break
			}
			line_ctr += 1
			if line_ctr > sa.page_len{
				page_ctr +=1
				line_ctr = 1
			}
			if (page_ctr >= sa.start_page) && (page_ctr <= sa.end_page){
				_, err := fout.Write([]byte(line))
				if err != nil{
					fmt.Println(err)
					os.Exit(1)
				}
			}
		}
	}else{
		page_ctr = 1
		for{
			page,  err := bufFin.ReadString('\f')
			if err != nil{
				break
			}
			if (page_ctr >= sa.start_page) && (page_ctr <= sa.end_page){
				_, err := fout.Write([]byte(page))
				if err != nil{
					os.Exit(1)
				}
			}
			page_ctr += 1
		}
	}

	if page_ctr < sa.start_page{
		fmt.Fprintf(os.Stderr, "\n%s: start_page (%d) greater than total pages (%d), "+
		" no output written\n", progname, sa.start_page, page_ctr)
	}else if page_ctr < sa.end_page{
		fmt.Fprintf(os.Stderr,"\n%s: end_page (%d) greater than total pages (%d),"+
		" less output than expected\n", progname, sa.end_page, page_ctr);
	}
}