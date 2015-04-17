#!/opt/sbin/dtrace -s

/*
 */

#pragma D option flowindent
this int op;
this int nhtype;
this int af;

inline string op_str =
    this->op == 1 ? "RTM_ADD":
    this->op == 2 ? "RTM_CHANGE":
    this->op == 3 ? "RTM_DELETE":
    this->op == 4 ? "RTM_GET":
    this->op == 5 ? "RTM_GETNEXT":
    this->op == 6 ? "RTM_CLEAR":
    this->op == 7 ? "RTM_DELETE_ALL":
    this->op == 8 ? "RTM_ADD_AND_GET":
    this->op == 9 ? "RTM_CHANGE_AND_GET":
    "UKNOWN";

inline string nhtype_str =
    this->nhtype == 0 ? "DISCARD" :
    this->nhtype == 1 ? "REJECT" :
    this->nhtype == 2 ? "UNICAST" :
    this->nhtype == 3 ? "UNILIST" :
    this->nhtype == 4 ? "IDXD" :
    this->nhtype == 5 ? "INDR" :
    this->nhtype == 6 ? "HOLD" :
    this->nhtype == 7 ? "RESOLV" :
    this->nhtype == 8 ? "LOCAL" :
    this->nhtype == 9 ? "RECEIVE" :
    this->nhtype == 10 ? "MULTIRT" :
    this->nhtype == 11 ? "BCAST" :
    this->nhtype == 12 ? "MCAST" :
    this->nhtype == 13 ? "MGROUP" :
    this->nhtype == 15 ? "TABLE" :
    this->nhtype == 16 ? "DENY" :
    this->nhtype == 19 ? "IFLIST" :
    this->nhtype == 24 ? "COMPOSITE" :
    "* NOT_YET_DEFINED *";

inline string af_str =
    this->af == 1 ? "AF_LOCAL" :
    this->af == 2 ? "AF_INET" :
    this->af == 7 ? "AF_ISO" :
    this->af == 28 ? "AF_INET6" :
    this->af == 37 ? "AF_TAG" :
    this->af == 38 ? "AF_TNP" :
    "* NOT_YET_DEFINED *";

inline string arg1str =
    arg1 == 0 ? "SUCCESS":
    arg1 == 1 ? "EPERM":
    arg1 == 2 ? "ENOENT":
    arg1 == 3 ? "ESRCH":
    arg1 == 4 ? "EINTR":
    arg1 == 5 ? "EIO":
    arg1 == 6 ? "ENXIO":
    arg1 == 7 ? "E2BIG":
    arg1 == 8 ? "ENOEXEC":
    arg1 == 9 ? "EBADF":
    arg1 == 10 ? "ECHILD":
    arg1 == 11 ? "EAGAIN":
    arg1 == 12 ? "ENOMEM":
    arg1 == 13 ? "EACCES":
    arg1 == 14 ? "EFAULT":
    arg1 == 15 ? "ENOTBLK":
    arg1 == 16 ? "EBUSY":
    arg1 == 17 ? "EEXIST":
    arg1 == 18 ? "EXDEV":
    arg1 == 19 ? "ENODEV":
    arg1 == 20 ? "ENOTDIR":
    arg1 == 21 ? "EISDIR":
    arg1 == 22 ? "EINVAL":
    arg1 == 23 ? "ENFILE":
    arg1 == 24 ? "EMFILE":
    arg1 == 25 ? "ENOTTY":
    arg1 == 26 ? "ETXTBSY":
    arg1 == 27 ? "EFBIG":
    arg1 == 28 ? "ENOSPC":
    arg1 == 29 ? "ESPIPE":
    arg1 == 30 ? "EROFS":
    arg1 == 31 ? "EMLINK":
    arg1 == 32 ? "EPIPE":
    arg1 == 33 ? "EDOM":
    arg1 == 34 ? "ERANGE":
    arg1 == 35 ? "ENOMSG":
    arg1 == 36 ? "EIDRM":
    arg1 == 37 ? "ECHRNG":
    arg1 == 38 ? "EL2NSYNC":
    arg1 == 39 ? "EL3HLT":
    arg1 == 40 ? "EL3RST":
    arg1 == 41 ? "ELNRNG":
    arg1 == 42 ? "EUNATCH":
    arg1 == 43 ? "ENOCSI":
    arg1 == 44 ? "EL2HLT":
    arg1 == 45 ? "EDEADLK":
    arg1 == 46 ? "ENOLCK":
    arg1 == 47 ? "ECANCELED":
    arg1 == 48 ? "ENOTSUP":
    arg1 == 49 ? "EDQUOT":
    arg1 == 50 ? "EBADE":
    arg1 == 51 ? "EBADR":
    arg1 == 52 ? "EXFULL":
    arg1 == 53 ? "ENOANO":
    arg1 == 54 ? "EBADRQC":
    arg1 == 55 ? "EBADSLT":
    arg1 == 56 ? "EDEADLOCK":
    arg1 == 57 ? "EBFONT":
    arg1 == 58 ? "EOWNERDEAD":
    arg1 == 59 ? "ENOTRECOVERABLE":
    arg1 == 60 ? "ENOSTR":
    arg1 == 61 ? "ENODATA":
    arg1 == 62 ? "ETIME":
    arg1 == 63 ? "ENOSR":
    arg1 == 64 ? "ENONET":
    arg1 == 65 ? "ENOPKG":
    arg1 == 66 ? "EREMOTE":
    arg1 == 67 ? "ENOLINK":
    arg1 == 68 ? "EADV":
    arg1 == 69 ? "ESRMNT":
    arg1 == 70 ? "ECOMM":
    arg1 == 71 ? "EPROTO":
    arg1 == 72 ? "ELOCKUNMAPPED":
    arg1 == 73 ? "ENOTACTIVE":
    arg1 == 74 ? "EMULTIHOP":
    arg1 == 77 ? "EBADMSG":
    arg1 == 78 ? "ENAMETOOLONG":
    arg1 == 79 ? "EOVERFLOW":
    arg1 == 80 ? "ENOTUNIQ":
    arg1 == 81 ? "EBADFD":
    arg1 == 82 ? "EREMCHG":
    arg1 == 83 ? "ELIBACC":
    arg1 == 84 ? "ELIBBAD":
    arg1 == 85 ? "ELIBSCN":
    arg1 == 86 ? "ELIBMAX":
    arg1 == 87 ? "ELIBEXEC":
    arg1 == 88 ? "EILSEQ":
    arg1 == 89 ? "ENOSYS":
    arg1 == 90 ? "ELOOP":
    arg1 == 91 ? "ERESTART":
    arg1 == 92 ? "ESTRPIPE":
    arg1 == 93 ? "ENOTEMPTY":
    arg1 == 94 ? "EUSERS":
    arg1 == 95 ? "ENOTSOCK":
    arg1 == 96 ? "EDESTADDRREQ":
    arg1 == 97 ? "EMSGSIZE":
    arg1 == 98 ? "EPROTOTYPE":
    arg1 == 99 ? "ENOPROTOOPT":
    arg1 == 120 ? "EPROTONOSUPPORT":
    arg1 == 121 ? "ESOCKTNOSUPPORT":
    arg1 == 122 ? "EOPNOTSUPP":
    arg1 == 123 ? "EPFNOSUPPORT":
    arg1 == 124 ? "EAFNOSUPPORT":
    arg1 == 125 ? "EADDRINUSE":
    arg1 == 126 ? "EADDRNOTAVAIL":
    arg1 == 127 ? "ENETDOWN":
    arg1 == 128 ? "ENETUNREACH":
    arg1 == 129 ? "ENETRESET":
    arg1 == 130 ? "ECONNABORTED":
    arg1 == 131 ? "ECONNRESET":
    arg1 == 132 ? "ENOBUFS":
    arg1 == 133 ? "EISCONN":
    arg1 == 134 ? "ENOTCONN":
    arg1 == 143 ? "ESHUTDOWN":
    arg1 == 144 ? "ETOOMANYREFS":
    arg1 == 145 ? "ETIMEDOUT":
    arg1 == 146 ? "ECONNREFUSED":
    arg1 == 147 ? "EHOSTDOWN":
    arg1 == 148 ? "EHOSTUNREACH":
    arg1 == 149 ? "EALREADY":
    arg1 == 150 ? "EINPROGRESS":
    arg1 == 151 ? "ESTALE":
        "UNKNOWN: Unknown error code.";

/*
 * route-table request functions
 */	
/*
fbt:kernel:rt_table_request:entry {
    this->op = args[0];
    this->af = ((rtb_msg_t *)args[1])->rtbm_af;
    printf("%s af: %s op: %s table-name:\n", 
	   execname, af_str, op_str); 
    stack(32);
}

fbt:kernel:tag_rttable_op:entry {
	this->op = args[1];
	this->nhtype = args[3];
	printf("rt-table-id %d op: %s nhtype: %s rt-flags: 0x%x\n",
	       ((struct route_table *)args[0])->rtb_id, op_str, nhtype_str, args[2]);
        stack(32);
}

fbt:kernel:rt_table_request:return {
}

fbt:kernel:rtb_rt_lookup:entry
/ args[0] != 0 /
{
    printf("%s rtb-name: %s\n", execname, ((struct route_table *)args[0])->rtb_name);
    stack(32);
}

fbt:kernel:rtb_rt_lookup:return
/ arg1 != 0/
{
	this->nhtype = ((rnh_t *)arg1)->rnh_type;
	printf("returned %p nhtype %s %d idx %d\n", arg1, nhtype_str, this->nhtype, 
	       ((rnh_t *)arg1)->rnh_index.y);
}
*/

/*
 * IP Forwarding. Bottom-Half 
 */

/*
 *  (a static function, can't insert probes)
fbt:kernel:_ip_NQ:entry{
        printf("%s ifqueue ptr  %p\n", execname, args[1]);
	stack(32);
    }
*/

/*
fbt:kernel:ip_recv_input_work:entry
/ args[0] != 0 /
{
    printf("%s rnhidx %d rcv-idx %d\n", execname,
           ((struct mbuf *)args[0])->M_dat.MH.MH_pkthdr.jp.jph_j_rnhidx.y, 
           ((struct mbuf *)args[0])->M_dat.MH.MH_pkthdr.jp.jph_j_rcvidx.x
	   );
	stack(32);
}
*/

/* RECEIVE NH output function */
/*
fbt:kernel:in_recv_output:entry
{
    stack(32);
}
*/

/* top-half input handler */
/*
fbt:kernel:ip_input:entry{
    this->nhtype = ((rnh_t *)args[0])->rnh_type;
    printf("%s nh-idx %d nh-type %s pkthdr.rnhidx %d ifl %d otype %d\n",
           execname, ((rnh_t *)args[0])->rnh_index.y, nhtype_str, 
           ((struct mbuf *)args[1])->M_dat.MH.MH_pkthdr.jp.jph_j_rnhidx.y, 
           ((struct mbuf *)args[1])->M_dat.MH.MH_pkthdr.jp.jph_j_rcvidx.x, 
           args[2]);
    stack(32);
}
*/

/* send from the same host */
/*
fbt:kernel:ip_local_input:entry{
    this->nhtype = ((rnh_t *)args[1])->rnh_type;
    printf("%s nh-idx %d nh-type %s pkthdr.rnhidx %d ifl %d\n",
           execname, ((rnh_t *)args[1])->rnh_index.y, nhtype_str, 
           ((struct mbuf *)args[0])->M_dat.MH.MH_pkthdr.jp.jph_j_rnhidx.y, 
           ((struct mbuf *)args[0])->M_dat.MH.MH_pkthdr.jp.jph_j_rcvidx.x);
    stack(32);
}
*/

